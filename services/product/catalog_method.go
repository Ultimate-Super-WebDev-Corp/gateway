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
		Size(int(msg.Limit))

	searchReq = applySorts(searchReq, msg)
	searchReq = applyFilters(searchReq, msg)
	searchReq = applyAggregations(searchReq)

	searchRes, err := searchReq.Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &product.CatalogResponse{
		Products:   buildProducts(ctx, searchRes),
		Filters:    buildFilters(ctx, searchRes),
		Categories: buildCategories(ctx, msg.CategoryId, searchRes),
		Sorts:      dictSorts,
		NextToken:  msg.Token + uint64(len(searchRes.Hits.Hits)),
	}, nil
}

func applySorts(searchReq *elastic.SearchService, msg *product.CatalogRequest) *elastic.SearchService {
	searchReq = searchReq.Sort(fieldScore, false)

	mapUniqueSorts := map[string]struct{}{}
	uniqueSorts := make([]string, 0, len(msg.Sorts))
	for _, s := range msg.Sorts {
		if _, ok := mapUniqueSorts[s.Id]; ok {
			continue
		}
		mapUniqueSorts[s.Id] = struct{}{}
		uniqueSorts = append(uniqueSorts, s.Id)
	}

	clearSorts := map[string]*product.Sort{}
	for _, s := range dictSorts {
		if _, ok := mapUniqueSorts[s.Id]; !ok {
			continue
		}
		clearSorts[s.Id] = s
	}

	for _, id := range uniqueSorts {
		cs, ok := clearSorts[id]
		if !ok {
			continue
		}
		searchReq = searchReq.Sort(cs.Id, cs.Ascending)
	}

	if len(clearSorts) == 0 {
		searchReq = searchReq.Sort(defaultSort.Id, defaultSort.Ascending)
	}

	searchReq = searchReq.Sort(fieldId, true)
	return searchReq
}

func applyFilters(searchReq *elastic.SearchService, msg *product.CatalogRequest) *elastic.SearchService {
	qMust := make([]elastic.Query, 0, len(msg.Filters)+1)
	qMust = append(qMust, elastic.MatchAllQuery{})
	uniqueFilters := map[string]*product.Filter{}
	for _, f := range msg.Filters {
		uniqueFilters[f.Id] = f
	}

	for _, f := range dictFilters {
		uf, ok := uniqueFilters[f.Id]
		if !ok {
			continue
		}
		switch v := uf.Value.(type) {
		case *product.Filter_ListFilter:
			qMust = append(
				qMust,
				elastic.NewTermsQuery(getEFilterField(uf.Id), stringArrayToInterfaceArray(v.ListFilter.List)...),
			)
		case *product.Filter_RangeFilter:
			qMust = append(
				qMust,
				elastic.NewRangeQuery(uf.Id).Lte(v.RangeFilter.Max).Gte(v.RangeFilter.Min),
			)
		}
	}

	if msg.CategoryId != "" {
		qMust = append(
			qMust,
			elastic.NewMatchQuery(getEFilterField(fieldCategories), msg.CategoryId),
		)
	}

	if msg.TextSearch != "" {
		qMust = append(qMust, makeTextSearchQuery(msg.TextSearch))
	}

	return searchReq.Query(elastic.NewBoolQuery().Must(qMust...))
}

func applyAggregations(searchReq *elastic.SearchService) *elastic.SearchService {
	for _, f := range dictFilters {
		switch f.Value.(type) {
		case *product.Filter_ListFilter:
			searchReq = searchReq.Aggregation(f.Id, elastic.NewTermsAggregation().Field(getEFilterField(f.Id)))
		case *product.Filter_RangeFilter:
			searchReq = searchReq.Aggregation(makeFilterRangeMaxName(f.Id), elastic.NewMaxAggregation().Field(getEFilterField(f.Id)))
			searchReq = searchReq.Aggregation(makeFilterRangeMinName(f.Id), elastic.NewMinAggregation().Field(getEFilterField(f.Id)))
		}
	}

	searchReq = searchReq.Aggregation(fieldCategories, elastic.NewTermsAggregation().Field(getEFilterField(fieldCategories)))
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
			aggRes, ok := searchRes.Aggregations.Filters(f.Id)
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Id))
				continue
			}
			list := make([]string, 0, len(aggRes.Buckets))
			for _, b := range aggRes.Buckets {
				list = append(list, b.Key.(string))
			}
			filters = append(filters, &product.Filter{
				Id:   f.Id,
				Name: f.Name,
				Value: &product.Filter_ListFilter{
					ListFilter: &product.ListFilter{
						List: list,
					},
				},
			})
		case *product.Filter_RangeFilter:
			aggResMax, ok := searchRes.Aggregations.Max(makeFilterRangeMaxName(f.Id))
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Id))
				continue
			}
			aggResMin, ok := searchRes.Aggregations.Min(makeFilterRangeMinName(f.Id))
			if !ok {
				ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", f.Id))
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
				Id:   f.Id,
				Name: f.Name,
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

func buildCategories(ctx context.Context, categoryId string, searchRes *elastic.SearchResult) []*product.Category {
	aggRes, ok := searchRes.Aggregations.Filters(fieldCategories)
	if !ok {
		ctxzap.Extract(ctx).Warn("categories not found")
		return []*product.Category{}
	}

	subtree := getSubtree(categoryId, dictCategoryTree)
	mapSubtree := map[string]struct{}{}
	for _, node := range subtree {
		mapSubtree[node.Id] = struct{}{}
	}

	mapCategories := map[string]struct{}{}
	for _, b := range aggRes.Buckets {
		id := b.Key.(string)
		if _, ok := mapSubtree[id]; !ok {
			continue
		}
		mapCategories[id] = struct{}{}
	}

	categories := make([]*product.Category, 0, len(mapCategories))
	for _, node := range subtree {
		if _, ok := mapCategories[node.Id]; !ok {
			continue
		}
		categories = append(categories, &product.Category{
			Id:   node.Id,
			Name: node.Name,
		})
	}
	return categories
}

func stringArrayToInterfaceArray(a []string) []interface{} {
	r := make([]interface{}, 0, len(a))
	for _, v := range a {
		r = append(r, v)
	}

	return r
}
