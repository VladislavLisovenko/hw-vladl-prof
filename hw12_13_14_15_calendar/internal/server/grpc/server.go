package internalgrpc

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/app"
	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	logg    app.Logger
	app     *app.App
	address string
}

func NewServer(logger app.Logger, app *app.App, host string, port string) *Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(loggerInterceptor(logger)))
	grpc.ChainStreamInterceptor()
	server := &Server{
		server:  grpcServer,
		logg:    logger,
		app:     app,
		address: net.JoinHostPort(host, port),
	}

	RegisterEventsServer(grpcServer, server)

	return server
}

func loggerInterceptor(logger app.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		start := time.Now()

		logger.Info(
			strings.Join(
				[]string{
					time.Now().String(),
					info.FullMethod,
					fmt.Sprintf("%s", req),
					time.Since(start).String(),
				},
				" ",
			),
		)

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func (s *Server) Start(ctx context.Context) error {
	lsn, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen an address: " + err.Error())
	}

	go func() {
		s.logg.Info(fmt.Sprintf("grpc server is running on %s", s.address))
		if err := s.server.Serve(lsn); err != nil {
			s.logg.Error("failed to start grpc server: " + err.Error())
		}
		s.logg.Info("grpc servers has been stopped")
	}()

	go func() {
		<-ctx.Done()
		s.server.Stop()
	}()

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
