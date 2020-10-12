package catalog

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/catalog"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Catalog) Search(ctx context.Context, msg *catalog.SearchRequest) (*catalog.SearchResponse, error) {
	searchRes, err := c.elasticCli.Search(objectProduct).
		Query(elastic.MatchAllQuery{}).
		From(int(msg.Token)).
		Size(int(msg.Limit)).
		Sort(catalog.OrderBy_name[int32(msg.Sort.OrderBy)], msg.Sort.Ascending).
		Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return &catalog.SearchResponse{
			Products:  []*catalog.Product{},
			NextToken: msg.Token,
		}, nil
	}

	return nil, nil
}
