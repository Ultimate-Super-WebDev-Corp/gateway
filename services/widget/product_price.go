package widget

import (
	"context"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/widget"
)

func (w Widget) ProductPrice(context.Context, *widget.ProductPriceRequest) (*widget.HtmlBody, error) {
	return &widget.HtmlBody{
		Body: []byte(""),
	}, nil
}
