package server

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewErrServer(code codes.Code, err error) ErrServer {
	return ErrServer{
		GrpcCode: code,
		Err:      err,
	}
}

type ErrServer struct {
	GrpcCode codes.Code
	Err      error
}

func (e ErrServer) GRPCStatus() *status.Status {
	return status.New(e.GrpcCode, e.Error())
}

func (e ErrServer) Error() string {
	return e.Err.Error()
}

func (e ErrServer) StackTrace() errors.StackTrace {
	if err, ok := e.Err.(interface{ StackTrace() errors.StackTrace }); ok {
		return err.StackTrace()
	}
	return nil
}
