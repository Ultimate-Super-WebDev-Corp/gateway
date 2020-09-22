package search

import (
	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

type Search struct {
	fileCli              file.FileClient
	imageAnnotatorClient *vision.ImageAnnotatorClient
}

type config struct{}

type Dependences struct {
	Registrar *grpc.Server
	FileCli   file.FileClient
}

func NewSearch(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	imageAnnotatorClient, err := vision.NewImageAnnotatorClient(context.Background())
	if err != nil {
		return errors.WithStack(err)
	}

	cus := &Search{
		fileCli:              dep.FileCli,
		imageAnnotatorClient: imageAnnotatorClient,
	}

	search.RegisterSearchServer(dep.Registrar, cus)

	return nil
}
