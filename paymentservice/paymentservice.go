package paymentservice

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jinzhu/copier"
	"github.com/jozuenoon/rest-api/paymentapi"
	"github.com/oklog/ulid"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const version = "v0.0.1"

var _ paymentapi.PaymentServiceServer = (*PaymentService)(nil)

func NewPaymentService(path string) (*PaymentService, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return &PaymentService{
		db:      db,
		entropy: entropy,
	}, nil
}

type PaymentService struct {
	db      *leveldb.DB
	entropy io.Reader
}

func (ps *PaymentService) GetPayment(ctx context.Context, req *paymentapi.PaymentRequest) (*paymentapi.PaymentServiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetPayment")
	defer span.Finish()
	if req.Id != "" {
		key, err := ParseUlid(req.Id)
		if err != nil {
			return nil, err
		}
		data, err := ps.db.Get(key, nil)
		if err != nil {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("payment does not exist %s", req.Id))
		}
		payment, err := unmarshalPayment(data)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &paymentapi.PaymentServiceResponse{
			Data: []*paymentapi.Payment{payment},
		}, nil
	}
	var payments []*paymentapi.Payment
	iter := ps.db.NewIterator(nil, nil)
	for iter.Next() {
		payment, err := unmarshalPayment(iter.Value())
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		payments = append(payments, payment)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentapi.PaymentServiceResponse{
		Data: payments,
	}, nil
}

func (ps *PaymentService) CreatePayment(ctx context.Context, pt *paymentapi.PaymentAttributes) (*paymentapi.PaymentServiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreatePayment")
	defer span.Finish()
	// Generate ULID for new resource.
	uid, err := ulid.New(ulid.Timestamp(time.Now()), ps.entropy)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	key, err := uid.MarshalBinary()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	organizationId, err := organizationIdFromContext(ctx)
	if err != nil {
		// should be always present from authorization middleware
		return nil, status.Error(codes.Internal, err.Error())
	}
	payment := &paymentapi.Payment{
		Id:             uid.String(),
		OrganisationId: organizationId,
		Type:           "Payment",
		Version:        version,
		Attributes:     pt,
	}
	data, err := payment.Marshal()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = ps.db.Put(key, data, nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentapi.PaymentServiceResponse{
		Data: []*paymentapi.Payment{payment},
	}, nil
}

func (ps *PaymentService) UpdatePayment(ctx context.Context, pt *paymentapi.PaymentUpdate) (*paymentapi.PaymentServiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdatePayment")
	defer span.Finish()
	key, err := ParseUlid(pt.Id)
	if err != nil {
		return nil, err
	}
	data, err := ps.db.Get(key, nil)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	baseMsg, err := unmarshalPayment(data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = copier.Copy(&baseMsg.Attributes, pt.PaymentAttributes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	data, err = baseMsg.Marshal()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = ps.db.Put(key, data, nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentapi.PaymentServiceResponse{
		Data: []*paymentapi.Payment{baseMsg},
	}, nil
}

func (ps *PaymentService) DeletePayment(ctx context.Context, pt *paymentapi.PaymentRequest) (*paymentapi.PaymentRequest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeletePayment")
	defer span.Finish()
	key, err := ParseUlid(pt.Id)
	if err != nil {
		return nil, err
	}
	err = ps.db.Delete(key, nil)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return pt, nil
}

func (ps *PaymentService) Close() error {
	return ps.db.Close()
}

func unmarshalPayment(data []byte) (*paymentapi.Payment, error) {
	payment := &paymentapi.Payment{}
	err := proto.Unmarshal(data, payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func organizationIdFromContext(ctx context.Context) (string, error) {
	// TODO: API Key should be injected at authorization middleware.
	return "example_organization", nil
}

func ParseUlid(s string) ([]byte, error) {
	u, err := ulid.Parse(s)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "provided ID is invalid")
	}
	return u.MarshalBinary()
}
