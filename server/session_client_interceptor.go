package server

import (
	"context"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s Server) UnarySessionClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	token, err := s.makeSessionToken(ctx)
	if err != nil {
		return err
	}

	outCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs(mdToken, token))

	var header metadata.MD
	opts = append(opts, grpc.Header(&header))
	err = invoker(outCtx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	session, err := s.getSession(header)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	SessionInCtxUpdate(ctx, session)
	return nil
}

func (s Server) StreamSessionClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	token, err := s.makeSessionToken(ctx)
	if err != nil {
		return nil, err
	}

	outCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs(mdToken, token))

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
			return status.Error(codes.Internal, err.Error())
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