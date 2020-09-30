package server

import (
	"context"
	"sync"

	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
	wrappedContext context.Context
	sendMsg        func(m interface{}) error
	once           *sync.Once
}

func (w wrappedServerStream) SendMsg(m interface{}) error {
	var err error
	w.once.Do(func() {
		err = w.sendMsg(m)
	})
	if err != nil {
		return err
	}
	return w.ServerStream.SendMsg(m)
}

func newWrappedServerStream(stream grpc.ServerStream) wrappedServerStream {
	if existing, ok := stream.(wrappedServerStream); ok {
		return existing
	}
	return wrappedServerStream{ServerStream: stream, wrappedContext: stream.Context(), once: &sync.Once{}}
}
