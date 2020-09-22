package search

import (
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

type Search struct {
	fileCli  file.FileClient
	visionRR *gcVisionRoundRobin
}

type config struct {
	GCVisionPathToKeys string `env:"GC_VISION_PATH_TO_KEYS"`
}

type Dependences struct {
	Registrar *grpc.Server
	FileCli   file.FileClient
}

func NewSearch(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	gcVisionPathToKeys := strings.Split(cfg.GCVisionPathToKeys, ";")

	visionRR, err := newGcVisionRoundRobin(gcVisionPathToKeys)
	if err != nil {
		return errors.WithStack(err)
	}

	cus := &Search{
		fileCli:  dep.FileCli,
		visionRR: visionRR,
	}

	search.RegisterSearchServer(dep.Registrar, cus)

	return nil
}
