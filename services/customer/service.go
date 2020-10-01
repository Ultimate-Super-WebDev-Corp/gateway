package customer

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

type Customer struct {
	gatewayDB squirrel.StatementBuilderType
}

type config struct {
	GatewayDB string `env:"GATEWAY_DB"`
}

type Dependences struct {
	Registrar *grpc.Server
	GrpcConn  *grpc.ClientConn
}

func NewCustomer(dep Dependences) (customer.CustomerClient, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	gatewayDB, err := sql.Open("postgres", cfg.GatewayDB)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := gatewayDB.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}

	cus := &Customer{
		gatewayDB: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(gatewayDB),
	}

	customer.RegisterCustomerServer(dep.Registrar, cus)

	return customer.NewCustomerClient(dep.GrpcConn), nil
}
