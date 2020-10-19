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
	searchReq := p.elasticCli.Search(objectProduct).
		From(int(msg.Token)).
		Size(int(msg.Limit)).
		Sort(fieldId, true)

	searchReq = applySorts(searchReq, msg)
	searchReq = applyFilters(searchReq, msg)
	searchReq = applyAggregations(searchReq)

	searchRes, err := searchReq.Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &product.CatalogResponse{
		Products:  buildProducts(ctx, searchRes),
		Filters:   buildFilters(ctx, searchRes),
		Sorts:     dictSorts,
		NextToken: msg.Token + uint64(len(searchRes.Hits.Hits)),
	}, nil
}

func applySorts(searchReq *elastic.SearchService, msg *product.CatalogRequest) *elastic.SearchService {
	uniqueSorts := map[string]*product.Sort{}
	for _, s := range msg.Sorts {
		uniqueSorts[s.Field] = s
	}

	for _, s := range dictSorts {
		us, ok := uniqueSorts[s.Field]
		if !ok {
			continue
		}
		searchReq = searchReq.Sort(us.Field, us.Ascending)
	}
	return searchReq
}

func applyFilters(searchReq *elastic.SearchService, msg *product.CatalogRequest) *elastic.SearchService {
	qMust := make([]elastic.Query, 0, len(msg.Filters)+1)
	qMust = append(qMust, elastic.MatchAllQuery{})
	uniqueFilters := map[string]*product.Filter{}
	for _, f := range msg.Filters {
		uniqueFilters[f.Field] = f
	}

	for _, f := range dictFilters {
		uf, ok := uniqueFilters[f.Field]
		if !ok {
			continue
		}
		switch v := uf.Value.(type) {
		case *product.Filter_ListFilter:
			qMust = append(
				qMust,
				elastic.NewTermsQuery(getEFilterField(uf.Field), stringArrayToInterfaceArray(v.ListFilter.List)...),
			)
		case *product.Filter_RangeFilter:
			qMust = append(
				qMust,
				elastic.NewRangeQuery(uf.Field).Lte(v.RangeFilter.Max).Gte(v.RangeFilter.Min),
			)
		}
	}

	return searchReq.Query(elastic.NewBoolQuery().Must(qMust...))
}

func applyAggregations(searchReq *elastic.SearchService) *elastic.SearchService {
	for _, f := range dictFilters {
		switch f.Value.(type) {
		case *product.Filter_ListFilter:
			searchReq = searchReq.Aggregation(f.Field, elastic.NewTermsAggregation().Field(getEFilterField(f.Field)))
		case *product.Filter_RangeFilter:
			searchReq = searchReq.Aggregation(makeFilterRangeMaxName(f.Field), elastic.NewMaxAggregation().Field(getEFilterField(f.Field)))
			searchReq = searchReq.Aggregation(makeFilterRangeMinName(f.Field), elastic.NewMinAggregation().Field(getEFilterField(f.Field)))
		}
	}
	return searchReq
}

func buildProducts(ctx context.Context, searchRes *elastic.SearchResult) []*product.CatalogProduct {
	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return []*product.CatalogProduct{}
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

	return products
}

func buildFilters(ctx context.Context, searchRes *elastic.SearchResult) []*product.Filter {
	filters := make([]*product.Filter, 0, len(dictFilters))
	for _, f := range dictFilters {
		switch f.Value.(type) {
		case *product.Filter_ListFilter:
			aggRes, ok := searchRes.Aggregations.Filters(f.Field)
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Field))
				continue
			}
			list := make([]string, 0, len(aggRes.Buckets))
			for _, b := range aggRes.Buckets {
				list = append(list, b.Key.(string))
			}
			filters = append(filters, &product.Filter{
				Field: f.Field,
				Name:  f.Name,
				Value: &product.Filter_ListFilter{
					ListFilter: &product.ListFilter{
						List: list,
					},
				},
			})
		case *product.Filter_RangeFilter:
			aggResMax, ok := searchRes.Aggregations.Max(makeFilterRangeMaxName(f.Field))
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Field))
				continue
			}
			aggResMin, ok := searchRes.Aggregations.Min(makeFilterRangeMinName(f.Field))
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Field))
				continue
			}
			if aggResMax.Value == nil {
				continue
			}
			if aggResMin.Value == nil {
				continue
			}
			max := int64(*aggResMax.Value)
			min := int64(*aggResMin.Value)

			filters = append(filters, &product.Filter{
				Field: f.Field,
				Name:  f.Name,
				Value: &product.Filter_RangeFilter{
					RangeFilter: &product.RangeFilter{
						Min: min,
						Max: max,
					},
				},
			})
		}
	}
	return filters
}

func stringArrayToInterfaceArray(a []string) []interface{} {
	r := make([]interface{}, 0, len(a))
	for _, v := range a {
		r = append(r, v)
	}

	return r
}
