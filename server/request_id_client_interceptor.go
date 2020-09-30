package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s Server) UnaryRequestIdClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestId := requestIdFromCtx(ctx)
	outCtx := metadata.AppendToOutgoingContext(ctx, mdRequestId, requestId)
	return invoker(outCtx, method, req, reply, cc, opts...)
}

func (s Server) StreamRequestIdClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	requestId := requestIdFromCtx(ctx)
	outCtx := metadata.AppendToOutgoingContext(ctx, mdRequestId, requestId)
	return streamer(outCtx, desc, cc, method, opts...)
}
