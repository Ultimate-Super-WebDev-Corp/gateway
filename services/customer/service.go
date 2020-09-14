package customer

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/server/servicesbp/cutomerpb"
)

type Customer struct {
	customerDB squirrel.StatementBuilderType
}

type config struct {
	CustomerDB string `env:"CUSTOMER_DB"`
}

type Dependences struct {
	Registrar grpc.ServiceRegistrar
}
func NewCustomer(dep Dependences)  error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	customerDB, err := sql.Open("postgres", cfg.CustomerDB)
	if err != nil {
		return  errors.WithStack(err)
	}

	if err := customerDB.Ping(); err != nil {
		return  errors.WithStack(err)
	}

	cus := &Customer{
		customerDB: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(customerDB),
	}

	cutomerpb.RegisterCustomerService(dep.Registrar, &cutomerpb.CustomerService{
		Create: cus.Create,
		Get:    cus.Get,
	})

	return nil
}
