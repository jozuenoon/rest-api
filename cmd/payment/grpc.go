package payment

import (
	"context"
	"net"

	"github.com/jozuenoon/rest-api/paymentapi"
	"github.com/jozuenoon/rest-api/paymentservice"
	"github.com/oklog/run"
	"google.golang.org/grpc"
)

func grpcServiceServer(_ context.Context, g *run.Group, databasePath string) error {
	ln, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return err
	}
	svc, err := paymentservice.NewPaymentService(databasePath)
	if err != nil {
		return err
	}
	g.Add(func() error {
		grpcServer := grpc.NewServer()
		paymentapi.RegisterPaymentServiceServer(grpcServer, svc)
		return grpcServer.Serve(ln)
	}, func(error) {
		ln.Close()
		svc.Close()
	})
	return nil
}
