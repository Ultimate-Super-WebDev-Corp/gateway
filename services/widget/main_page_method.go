package widget

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/widget"
)

func (w Widget) MainPage(context.Context, *empty.Empty) (*widget.HtmlBody, error) {
	return &widget.HtmlBody{
		Body: []byte(""),
	}, nil
}
