package customer

import (
	"context"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Create(ctx context.Context, msg *customer.CreateRequest) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId != 0 {
		return nil, status.Error(codes.PermissionDenied, "the session has a customer")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(msg.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
				return nil, status.Error(codes.AlreadyExists, "customer already exists")
			}
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	session.CustomerId = customerId
	return &customer.CustomerMsg{
		Email: msg.Customer.Email,
		Name:  msg.Customer.Name,
	}, nil
}
