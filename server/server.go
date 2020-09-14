package server

import (
	"net"

	"github.com/caarlos0/env/v6"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//todo make pkg
type config struct {
	Port string `env:"PORT" envDefault:":8080"`
}

type Server struct {
	RpcServer *grpc.Server
	Logger    *zap.Logger

	cfg config
}

func NewServer() Server {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(errors.WithStack(err))
	}

	logger := initZapLog()

	s := grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			grpcZap.StreamServerInterceptor(logger),
			grpcValidator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			grpcZap.UnaryServerInterceptor(logger),
			grpcValidator.UnaryServerInterceptor(),
		)),
	)

	return Server{
		RpcServer: s,
		Logger:    logger,
		cfg:       cfg,
	}
}

func (s Server) Serve() error {
	defer func() {
		if err := s.Logger.Sync(); err != nil {
			s.Logger.Error("logger sync error", zap.Any("error", errors.WithStack(err)))
		}
	}()
	lis, err := net.Listen("tcp", s.cfg.Port)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := s.RpcServer.Serve(lis); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
