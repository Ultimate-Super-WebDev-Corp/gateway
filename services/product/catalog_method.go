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

	searchReq = applySorts(msg, searchReq)
	searchReq = applyFilters(ctx, msg, searchReq)

	searchRes, err := searchReq.Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &product.CatalogResponse{
		Products:  buildProducts(ctx, searchRes),
		NextToken: buildNextToken(msg, searchRes),
		Total:     buildTotal(searchRes),
	}, nil
}

func applySorts(msg *product.CatalogRequest, searchReq *elastic.SearchService) *elastic.SearchService {
	if msg.TextSearch != "" {
		searchReq = searchReq.Sort(fieldScore, false)
	}

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

func applyFilters(ctx context.Context, msg *product.CatalogRequest, searchReq *elastic.SearchService) *elastic.SearchService {
	eFilters, eMust := buildEFiltersAndEMust(ctx, msg.Filters, msg.SelectedCategoryId, msg.TextSearch)
	return searchReq.Query(elastic.NewBoolQuery().Filter(eFilters...).Must(eMust...))
}

func buildEFiltersAndEMust(ctx context.Context, filters []*product.Filter, selectedCategoryId string, textSearch string, excludeFilterIds ...string) (eFilters []elastic.Query, eMust []elastic.Query) {
	eMust = make([]elastic.Query, 0, 1)
	eFilters = make([]elastic.Query, 0, len(filters)+1)
	eFilters = append(eFilters, elastic.MatchAllQuery{})

	mapExcludeFilterIds := map[string]struct{}{}
	for _, id := range excludeFilterIds {
		mapExcludeFilterIds[id] = struct{}{}
	}

	uniqueFilters := map[string]*product.Filter{}
	for _, f := range filters {
		if _, ok := mapExcludeFilterIds[f.Id]; ok {
			continue
		}
		uniqueFilters[f.Id] = f
	}

	for _, f := range dictFilters {
		uf, ok := uniqueFilters[f.Id]
		if !ok {
			continue
		}
		switch v := uf.Value.(type) {
		case *product.Filter_ListFilter:
			eFilters = append(
				eFilters,
				elastic.NewTermsQuery(getEFilterField(uf.Id), stringArrayToInterfaceArray(v.ListFilter.SelectedItems)...),
			)
		case *product.Filter_RangeFilter:
			eFilters = append(
				eFilters,
				elastic.NewRangeQuery(uf.Id).Lte(v.RangeFilter.SelectedValue.Max).Gte(v.RangeFilter.SelectedValue.Min),
			)
		case *product.Filter_SwitchFilter:
			eFilters = append(
				eFilters,
				dictSwitchFilter.getEQuery(uf.Id, v.SwitchFilter.SelectedSwitch),
			)
		default:
			ctxzap.Extract(ctx).Warn("unknown filter type", zap.String("filter", f.Id))
			continue
		}
	}

	if selectedCategoryId != "" {
		eFilters = append(
			eFilters,
			elastic.NewMatchQuery(getEFilterField(fieldCategories), selectedCategoryId),
		)
	}

	if textSearch != "" {
		eMust = append(eMust, makeTextSearchQuery(textSearch))
	}

	return eFilters, eMust
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

func buildTotal(searchRes *elastic.SearchResult) uint64 {
	if searchRes == nil || searchRes.Hits == nil || searchRes.Hits.TotalHits == nil {
		return 0
	}
	return uint64(searchRes.Hits.TotalHits.Value)
}

func buildNextToken(msg *product.CatalogRequest, searchRes *elastic.SearchResult) uint64 {
	if searchRes == nil || searchRes.Hits == nil {
		return msg.Token
	}
	return msg.Token + uint64(len(searchRes.Hits.Hits))
}

func stringArrayToInterfaceArray(a []string) []interface{} {
	r := make([]interface{}, 0, len(a))
	for _, v := range a {
		r = append(r, v)
	}

	return r
}
