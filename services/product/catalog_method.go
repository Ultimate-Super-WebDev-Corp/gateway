package product

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) Catalog(ctx context.Context, msg *product.CatalogRequest) (*product.CatalogResponse, error) {
	qMust := make([]elastic.Query, 0, len(msg.Filter)+1)
	qMust = append(qMust, elastic.MatchAllQuery{})
	for _, f := range msg.Filter {
		qMust = append(
			qMust,
			elastic.NewTermsQuery(getEFilterField(f.Field), stringArrayToInterfaceArray(f.List)...),
		)
	}

	if msg.Sort == nil {
		msg.Sort = &product.Sort{}
	}
	searchReq := p.elasticCli.Search(objectProduct).
		Query(elastic.NewBoolQuery().Must(qMust...)).
		From(int(msg.Token)).
		Size(int(msg.Limit)).
		Sort(product.OrderBy_name[int32(msg.Sort.OrderBy)], msg.Sort.Ascending).
		Sort(fieldId, true)

	for _, f := range filtersList {
		searchReq = searchReq.Aggregation(f.name, f.aggregation)
	}

	searchRes, err := searchReq.Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return &product.CatalogResponse{
			Products:  []*product.CatalogProduct{},
			NextToken: msg.Token,
		}, nil
	}

	products := make([]*product.CatalogProduct, 0, len(searchRes.Hits.Hits))
	for _, h := range searchRes.Hits.Hits {
		p := product.CatalogProduct{}
		if err := json.Unmarshal(h.Source, &p); err != nil {
			ctxzap.Extract(ctx).Error("invalid product", zap.Error(err))
			continue
		}
		products = append(products, &p)
	}

	filters := make([]*product.Filter, 0, len(filtersList))
	for _, f := range filtersList {
		aggRes, ok := searchRes.Aggregations.Filters(f.name)
		if !ok {
			continue
		}
		list := make([]string, 0, len(aggRes.Buckets))
		for _, b := range aggRes.Buckets {
			list = append(list, b.Key.(string))
		}
		filters = append(filters, &product.Filter{
			Field: f.name,
			List:  list,
		})
	}

	return &product.CatalogResponse{
		Products:  products,
		Filter:    filters,
		NextToken: msg.Token + uint64(len(searchRes.Hits.Hits)),
	}, nil
}

func stringArrayToInterfaceArray(a []string) []interface{} {
	r := make([]interface{}, 0, len(a))
	for _, v := range a {
		r = append(r, v)
	}

	return r
}
