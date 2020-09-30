package server

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const mdRequestId = "request_id"

func (s Server) UnaryRequestIdServerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	requestId, err := s.getRequestId(md)
	if err != nil {
		return nil, NewErrServer(codes.Internal, errors.WithStack(err))
	}
	ctx = requestIdToCtx(ctx, requestId)

	return handler(ctx, req)
}

func (s Server) StreamRequestIdServerInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrapped := newWrappedServerStream(ss)
	md, _ := metadata.FromIncomingContext(wrapped.wrappedContext)
	requestId, err := s.getRequestId(md)
	if err != nil {
		return NewErrServer(codes.Internal, errors.WithStack(err))
	}

	wrapped.wrappedContext = requestIdToCtx(wrapped.wrappedContext, requestId)
	return handler(srv, wrapped)
}
