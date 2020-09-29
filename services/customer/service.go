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
	GCVisionPathToKeys string `env:"GC_VISION_PATH_TO_KEYS"`
	ElasticUrls        string `env:"ELASTIC_URLS"`
	GatewayDB          string `env:"GATEWAY_DB"`
}

type Dependences struct {
	Registrar *grpc.Server
}

func NewCustomer(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	gatewayDB, err := sql.Open("postgres", cfg.GatewayDB)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := gatewayDB.Ping(); err != nil {
		return errors.WithStack(err)
	}

	pr := &Customer{
		gatewayDB: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(gatewayDB),
	}

	customer.RegisterCustomerServer(dep.Registrar, pr)

	return nil
}
