package customer

import (
	"context"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Create(ctx context.Context, msg *customer.CreateRequest) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId != 0 {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("the session has a customer"))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(msg.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	query := c.gatewayDB.Insert(objectCustomer).
		Columns(fieldEmail, fieldPassword, fieldName).
		Values(msg.Customer.Email, string(password), msg.Customer.Name).
		Suffix("returning id").
		QueryRow()

	customerId := int64(0)
	if err := query.Scan(&customerId); err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return nil, server.NewErrServer(codes.AlreadyExists, errors.New("customer already exists"))
			}
		}

		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	session.CustomerId = customerId
	return &customer.CustomerMsg{
		Email: msg.Customer.Email,
		Name:  msg.Customer.Name,
	}, nil
}
