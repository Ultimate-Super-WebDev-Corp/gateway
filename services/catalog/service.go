package catalog

import (
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/catalog"
)

type Catalog struct {
	elasticCli *elastic.Client
}

type config struct {
	ElasticUrls string `env:"ELASTIC_URLS"`
}

type Dependences struct {
	Registrar *grpc.Server
}

func NewCatalog(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
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

	c := &Catalog{
		elasticCli: elasticCli,
	}

	catalog.RegisterCatalogServer(dep.Registrar, c)
	return nil
}
