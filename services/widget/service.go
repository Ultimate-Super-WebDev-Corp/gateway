package widget

import (
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/widget"
)

type Widget struct{}

type Dependences struct {
	Registrar *grpc.Server
}

func NewWidget(dep Dependences) error {
	w := &Widget{}

	widget.RegisterWidgetServer(dep.Registrar, w)

	return nil
}
