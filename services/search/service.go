package search

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

type Search struct {
	fileSrv file.FileServer
}

type config struct{}

type Dependences struct {
	Registrar *grpc.Server
	FileSrv   file.FileServer
}

func NewSearch(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	cus := &Search{
		fileSrv: dep.FileSrv,
	}

	search.RegisterSearchServer(dep.Registrar, cus)

	return nil
}
