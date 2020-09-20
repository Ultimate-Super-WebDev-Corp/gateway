package search

import (
	"context"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

func (s Search) SearchByIds(ctx context.Context, msg *search.SearchByIdsRequest) (*search.Product, error) {
	return &search.Product{}, nil
}
