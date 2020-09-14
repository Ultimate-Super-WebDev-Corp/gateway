package customer

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/servicesbp/cutomerpb"
)

func (c *Customer) Get(ctx context.Context, req *cutomerpb.CustomerGetRequest) (*cutomerpb.CustomerResponse, error) {
	logger := ctxzap.Extract(ctx)

	resp := &cutomerpb.CustomerResponse{}
	row := c.customerDB.Select("id", "email", "name").
		From(customerTableName).
		Where(squirrel.Eq{
			"email":    req.Email,
			"password": req.Password,
		}).QueryRow()

	if err := row.Scan(&resp.Id, &resp.Email, &resp.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}

		logger.Error("customer select error", zap.Any("error", errors.WithStack(err)))
		return nil, status.Error(codes.Internal, "customer get error")
	}

	return resp, nil
}
