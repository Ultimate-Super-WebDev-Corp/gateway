package server

import (
	"context"
	"io"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func (s Server) UnarySessionClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	token, err := s.makeSessionToken(ctx)
	if err != nil {
		return err
	}

	outCtx := metadata.AppendToOutgoingContext(ctx, mdToken, token)

	var header metadata.MD
	opts = append(opts, grpc.Header(&header))

	defer func() {
		session, err := s.getSession(header)
		if err != nil {
			return
		}
		SessionInCtxUpdate(ctx, session)
	}()

	return invoker(outCtx, method, req, reply, cc, opts...)
}

func (s Server) StreamSessionClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	token, err := s.makeSessionToken(ctx)
	if err != nil {
		return nil, err
	}

	outCtx := metadata.AppendToOutgoingContext(ctx, mdToken, token)

	var header metadata.MD
	opts = append(opts, grpc.Header(&header))
	resp, err := streamer(outCtx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	newResp := newWrappedSessionClientStream(resp)
	newResp.wrappedContext = ctx
	newResp.recvMsg = func(_ interface{}) error {
		session, err := s.getSession(header)
		if err != nil {
			return NewErrServer(codes.Internal, errors.WithStack(err))
		}
		SessionInCtxUpdate(ctx, session)
		return nil
	}

	return newResp, nil
}

type wrappedSessionClientStream struct {
	grpc.ClientStream
	wrappedContext context.Context
	recvMsg        func(m interface{}) error
	once           *sync.Once
}

func (w wrappedSessionClientStream) RecvMsg(m interface{}) error {
	err := w.ClientStream.RecvMsg(m)
	if err != nil && err != io.EOF {
		return err
	}
	if err == io.EOF {
		w.once.Do(func() {
			if wErr := w.recvMsg(m); wErr != nil {
				err = wErr
			}
		})
	}

	return err
}

func newWrappedSessionClientStream(stream grpc.ClientStream) wrappedSessionClientStream {
	if existing, ok := stream.(wrappedSessionClientStream); ok {
		return existing
	}
	return wrappedSessionClientStream{ClientStream: stream, wrappedContext: stream.Context(), once: &sync.Once{}}
}
