package review

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
)

type Review struct {
	customerCli      customer.CustomerClient
	gatewayDB        *sql.DB
	statementBuilder squirrel.StatementBuilderType
}

type config struct {
	GatewayDB string `env:"GATEWAY_DB"`
}

type Dependences struct {
	Registrar   *grpc.Server
	CustomerCli customer.CustomerClient
}

func NewReview(dep Dependences) error {
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

	c := &Review{
		gatewayDB:        gatewayDB,
		statementBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		customerCli:      dep.CustomerCli,
	}

	review.RegisterReviewServer(dep.Registrar, c)

	return nil
}
