package server

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const mdToken = "token"

func (s Server) UnarySessionServerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NewErrServer(codes.Unauthenticated, errors.New("metadata is not provided"))
	}

	session, err := s.getSession(md)
	if err != nil {
		return nil, NewErrServer(codes.Unauthenticated, errors.WithStack(err))
	}

	ctx = sessionToCtx(ctx, session)
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	respToken, err := s.makeSessionToken(ctx)
	if err != nil {
		return nil, NewErrServer(codes.Internal, errors.WithStack(err))
	}

	err = grpc.SendHeader(ctx, metadata.Pairs(mdToken, respToken))
	if err != nil {
		return nil, NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return resp, nil
}

func (s Server) StreamSessionServerInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return NewErrServer(codes.Unauthenticated, errors.New("metadata is not provided"))
	}

	session, err := s.getSession(md)
	if err != nil {
		return NewErrServer(codes.Unauthenticated, errors.WithStack(err))
	}

	wrapped := newWrappedSessionServerStream(ss)
	ctx = sessionToCtx(ctx, session)
	wrapped.wrappedContext = ctx
	wrapped.sendMsg = func(_ interface{}) error {
		respToken, err := s.makeSessionToken(ctx)
		if err != nil {
			return NewErrServer(codes.Internal, errors.WithStack(err))
		}

		err = ss.SetHeader(metadata.Pairs(mdToken, respToken))
		if err != nil {
			return NewErrServer(codes.Internal, errors.WithStack(err))
		}
		return nil
	}

	return handler(srv, wrapped)
}

type wrappedSessionServerStream struct {
	grpc.ServerStream
	wrappedContext context.Context
	sendMsg        func(m interface{}) error
	once           *sync.Once
}

func (w wrappedSessionServerStream) SendMsg(m interface{}) error {
	var err error
	w.once.Do(func() {
		err = w.sendMsg(m)
	})
	if err != nil {
		return err
	}
	return w.ServerStream.SendMsg(m)
}

func newWrappedSessionServerStream(stream grpc.ServerStream) wrappedSessionServerStream {
	if existing, ok := stream.(wrappedSessionServerStream); ok {
		return existing
	}
	return wrappedSessionServerStream{ServerStream: stream, wrappedContext: stream.Context(), once: &sync.Once{}}
}
