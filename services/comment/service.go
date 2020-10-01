package comment

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/comment"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

type Comment struct {
	gatewayDB   squirrel.StatementBuilderType
	customerCli customer.CustomerClient
}

type config struct {
	GatewayDB string `env:"GATEWAY_DB"`
}

type Dependences struct {
	Registrar   *grpc.Server
	CustomerCli customer.CustomerClient
}

func NewComment(dep Dependences) error {
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

	c := &Comment{
		gatewayDB:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(gatewayDB),
		customerCli: dep.CustomerCli,
	}

	comment.RegisterCommentServer(dep.Registrar, c)

	return nil
}
