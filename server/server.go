package server

import (
	"context"
	"net"
	"runtime/debug"

	"github.com/caarlos0/env/v6"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type config struct {
	IsDev     bool   `env:"IS_DEV"`
	Port      string `env:"PORT" envDefault:":8080"`
	SecretKey string `env:"TOKEN" envDefault:"TEST"`
}

type Server struct {
	RpcServer *grpc.Server
	Logger    *zap.Logger

	cfg config
}

func NewServer() (*Server, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	logger := initZapLog(cfg)
	logger.Sugar().Infof("starting http server on %s", cfg.Port)

	srv := &Server{
		Logger: logger,
		cfg:    cfg,
	}

	srv.RpcServer = grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			grpcZap.StreamServerInterceptor(logger, grpcZap.WithMessageProducer(srv.messageFunc)),
			grpcValidator.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandler(srv.recover)),
			srv.StreamRequestIdServerInterceptor,
			srv.StreamSessionServerInterceptor,
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			grpcZap.UnaryServerInterceptor(logger, grpcZap.WithMessageProducer(srv.messageFunc)),
			grpcValidator.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(srv.recover)),
			srv.UnaryRequestIdServerInterceptor,
			srv.UnarySessionServerInterceptor,
		)),
	)

	return srv, nil
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

func (s Server) GrpcDial() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		"localhost"+s.cfg.Port, grpc.WithInsecure(),
		grpc.WithStreamInterceptor(
			middleware.ChainStreamClient(
				s.StreamRequestIdClientInterceptor,
				s.StreamSessionClientInterceptor,
			)),
		grpc.WithUnaryInterceptor(
			middleware.ChainUnaryClient(
				s.UnaryRequestIdClientInterceptor,
				s.UnarySessionClientInterceptor,
			)),
	)
	return conn, errors.WithStack(err)
}

func (s Server) recover(p interface{}) (err error) {
	s.Logger.Sugar().Errorf("panic triggered: %v stacktrace from panic: %s", p, string(debug.Stack()))
	return status.Errorf(codes.Internal, "panic triggered: %v", p)
}

func (s Server) messageFunc(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
	zapFields := make([]zap.Field, 0)
	zapFields = append(zapFields, zap.Error(err))
	zapFields = append(zapFields, zap.String("grpc.code", code.String()))
	zapFields = append(zapFields, duration)
	if err, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		zapFields = append(zapFields, zap.Any("stacktrace", err.StackTrace()))
	}

	ctxzap.Extract(ctx).Check(level, msg).Write(zapFields...)
}
