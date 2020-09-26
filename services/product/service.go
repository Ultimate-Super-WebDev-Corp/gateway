package product

import (
	"database/sql"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

type Product struct {
	fileCli    file.FileClient
	visionRR   *gcVisionRoundRobin
	elasticCli *elastic.Client
	gatewayDB  squirrel.StatementBuilderType
}

type config struct {
	GCVisionPathToKeys string `env:"GC_VISION_PATH_TO_KEYS"`
	ElasticUrls        string `env:"ELASTIC_URLS"`
	GatewayDB          string `env:"GATEWAY_DB"`
}

type Dependences struct {
	Registrar *grpc.Server
	FileCli   file.FileClient
}

func NewProduct(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	visionRR, err := newGcVisionRoundRobin(cfg.GCVisionPathToKeys)
	if err != nil {
		return errors.WithStack(err)
	}

	elasticURLs := strings.Split(cfg.ElasticUrls, ";")
	elasticCli, err := elastic.NewClient(
		elastic.SetURL(elasticURLs...),
		elastic.SetSniff(false),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	gatewayDB, err := sql.Open("postgres", cfg.GatewayDB)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := gatewayDB.Ping(); err != nil {
		return errors.WithStack(err)
	}

	pr := &Product{
		fileCli:    dep.FileCli,
		visionRR:   visionRR,
		elasticCli: elasticCli,
		gatewayDB:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(gatewayDB),
	}

	product.RegisterProductServer(dep.Registrar, pr)

	return nil
}
