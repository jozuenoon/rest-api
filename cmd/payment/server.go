package payment

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/oklog/run"
)

type Server struct {
	DatabasePath string
}

func (s *Server) Run(ctx context.Context) {
	g := &run.Group{}
	cctx, cancel := context.WithCancel(ctx)

	// Since we pass context deeper inside services
	// we need to watch context Done and cancel in case of termination.
	g.Add(func() error {
		<-cctx.Done()
		return fmt.Errorf("context canceled")
	}, func(error) {
		cancel()
	})

	defer glog.Flush()

	signalHandler(cctx, g)
	err := httpPaymentGateway(cctx, g)
	if err != nil {
		glog.Errorln("msg=", "failed to create http server ", "err=", err)
	}

	err = grpcServiceServer(cctx, g, s.DatabasePath)
	if err != nil {
		glog.Errorln("msg", "failed to create grpc server ", "err=", err)
	}
	glog.Error("err", g.Run())
}

var shutdownSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func signalHandler(_ context.Context, g *run.Group) {
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, shutdownSignals...)
	g.Add(func() error {
		for sig := range signals {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				glog.Info("msg=", "got termination signal ", "sig=", sig)
				return fmt.Errorf("terminated: %v", sig)
			}
		}
		return nil
	}, func(error) {})
}
