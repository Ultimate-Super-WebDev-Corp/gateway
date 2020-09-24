package product

import (
	"strings"

	"github.com/caarlos0/env/v6"
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
}

type config struct {
	GCVisionPathToKeys string `env:"GC_VISION_PATH_TO_KEYS"`
	ElasticUrls        string `env:"ELASTIC_URLS"`
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

	gcVisionPathToKeys := strings.Split(cfg.GCVisionPathToKeys, ";")
	visionRR, err := newGcVisionRoundRobin(gcVisionPathToKeys)
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

	pr := &Product{
		fileCli:    dep.FileCli,
		visionRR:   visionRR,
		elasticCli: elasticCli,
	}

	product.RegisterProductServer(dep.Registrar, pr)

	return nil
}
