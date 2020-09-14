package customer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/servicesbp/cutomerpb"
)

const customerTableName = "customer"

func (c *Customer) Create(ctx context.Context, req *cutomerpb.CustomerCreateRequest) (*empty.Empty, error) {
	logger := ctxzap.Extract(ctx)

	_, err := c.customerDB.Insert(customerTableName).
		Columns("email", "password", "name").
		Values(req.Email, req.Password, req.Name).Exec()

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return nil, status.Error(codes.AlreadyExists, "customer already exists")
			}
		}

		logger.Error("customer insert error", zap.Any("error", errors.WithStack(err)))
		return nil, status.Error(codes.Internal, "customer create error")
	}

	return &empty.Empty{}, nil
}
