package payment

import (
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/gorilla/handlers"
	gw "github.com/jozuenoon/rest-api/paymentapi"
	"github.com/oklog/run"
)

func httpPaymentGateway(ctx context.Context, g *run.Group) error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	g.Add(func() error {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := gw.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
		if err != nil {
			return err
		}
		tracer, closer, err := makeJaegerTracer()
		if err != nil {
			return err
		}
		opentracing.SetGlobalTracer(tracer)
		defer closer.Close()
		h := tracingWrapper(mux)
		h = handlers.LoggingHandler(os.Stdout, h)
		h = handlers.CORS(handlers.AllowCredentials(),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-API-Key", "X-Client-Session-Id", "x-requested-with", "Content-Type", "Authorization"}))(h)
		h = handlers.ContentTypeHandler(h, "application/json", "application/vnd.api+json")
		return http.Serve(ln, h)
	}, func(error) {
		ln.Close()
	})
	return nil
}
