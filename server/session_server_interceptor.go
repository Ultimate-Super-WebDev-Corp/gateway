package server

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const mdToken = "token"

func (s Server) UnarySessionServerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NewErrServer(codes.Unauthenticated, errors.New("metadata is not provided"))
	}

	session, err := s.getSession(md)
	if err != nil {
		return nil, NewErrServer(codes.Unauthenticated, errors.WithStack(err))
	}
	ctx = sessionToCtx(ctx, session)

	defer func() {
		respToken, err := s.makeSessionToken(ctx)
		if err != nil {
			return
		}

		err = grpc.SendHeader(ctx, metadata.Pairs(mdToken, respToken))
		if err != nil {
			return
		}
	}()

	return handler(ctx, req)
}

func (s Server) StreamSessionServerInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	wrapped := newWrappedServerStream(ss)
	md, ok := metadata.FromIncomingContext(wrapped.wrappedContext)
	if !ok {
		return NewErrServer(codes.Unauthenticated, errors.New("metadata is not provided"))
	}

	session, err := s.getSession(md)
	if err != nil {
		return NewErrServer(codes.Unauthenticated, errors.WithStack(err))
	}

	wrapped.wrappedContext = sessionToCtx(wrapped.wrappedContext, session)
	wrapped.sendMsg = func() error {
		respToken, err := s.makeSessionToken(wrapped.wrappedContext)
		if err != nil {
			return NewErrServer(codes.Internal, errors.WithStack(err))
		}

		err = ss.SetHeader(metadata.Pairs(mdToken, respToken))
		if err != nil {
			return NewErrServer(codes.Internal, errors.WithStack(err))
		}
		return nil
	}

	defer func() {
		wrapped.once.Do(func() {
			err = wrapped.sendMsg()
		})
	}()
	return handler(srv, wrapped)
}
