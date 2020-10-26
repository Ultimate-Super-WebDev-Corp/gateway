package product

import (
	"context"
	"sort"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) CatalogMeta(ctx context.Context, msg *product.CatalogMetaRequest) (*product.CatalogMetaResponse, error) {
	searchReq := p.elasticCli.Search(objectProduct).
		From(0).
		Size(0)

	searchReq = applyAggregations(ctx, msg, searchReq)

	searchRes, err := searchReq.Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &product.CatalogMetaResponse{
		Filters:    buildFilters(ctx, msg.Filters, searchRes),
		Categories: buildCategories(ctx, msg.SelectedCategoryId, searchRes),
		Sorts:      dictSorts,
	}, nil
}

func applyAggregations(ctx context.Context, msg *product.CatalogMetaRequest, searchReq *elastic.SearchService) *elastic.SearchService {
	for _, f := range dictFilters {
		eFilters, eMust := buildEFiltersAndEMust(ctx, msg.Filters, msg.SelectedCategoryId, msg.TextSearch, f.Id)
		agg := elastic.NewFiltersAggregation().Filter(elastic.NewBoolQuery().Filter(eFilters...).Filter(eMust...))
		switch f.Value.(type) {
		case *product.Filter_ListFilter:
			agg = agg.SubAggregation(f.Id, elastic.NewTermsAggregation().Field(getEFilterField(f.Id)).Missing(dictAggregationMissingValue[f.Id]))
		case *product.Filter_RangeFilter:
			agg = agg.SubAggregation(makeFilterRangeMaxName(f.Id), elastic.NewMaxAggregation().Field(getEFilterField(f.Id)).Missing(dictAggregationMissingValue[f.Id]))
			agg = agg.SubAggregation(makeFilterRangeMinName(f.Id), elastic.NewMinAggregation().Field(getEFilterField(f.Id)).Missing(dictAggregationMissingValue[f.Id]))
		case *product.Filter_SwitchFilter:
			agg = agg.SubAggregation(f.Id, elastic.NewTermsAggregation().Field(getEFilterField(f.Id)).Missing(dictAggregationMissingValue[f.Id]))
		default:
			ctxzap.Extract(ctx).Warn("unknown filter type", zap.String("filter", f.Id))
			continue
		}

		searchReq = searchReq.Aggregation(f.Id, agg)
	}

	eFilters, eMust := buildEFiltersAndEMust(ctx, msg.Filters, msg.SelectedCategoryId, msg.TextSearch)

	return searchReq.Aggregation(fieldCategories, elastic.NewFiltersAggregation().Filter(elastic.NewBoolQuery().Filter(eFilters...).Filter(eMust...)).
		SubAggregation(fieldCategories, elastic.NewTermsAggregation().Field(getEFilterField(fieldCategories))))
}

func buildFilters(ctx context.Context, filters []*product.Filter, searchRes *elastic.SearchResult) []*product.Filter {
	selectedListFilters, selectedRangeFilters, selectedSwitchFilters := buildMapSelectedFilters(ctx, filters)

	respFilters := make([]*product.Filter, 0, len(dictFilters))
	for _, f := range dictFilters {
		switch f.Value.(type) {
		case *product.Filter_ListFilter:
			availableItems := buildListFilterAvailableItems(ctx, f.Id, searchRes)
			selectedItems := selectedListFilters[f.Id]
			if len(availableItems) == 0 && len(selectedItems) == 0 {
				continue
			}

			respFilters = append(respFilters, &product.Filter{
				Id:   f.Id,
				Name: f.Name,
				Value: &product.Filter_ListFilter{
					ListFilter: &product.ListFilter{
						AvailableItems: availableItems,
						SelectedItems:  selectedItems,
					},
				},
			})
		case *product.Filter_RangeFilter:
			availableValue := buildRangeFilterAvailableValue(ctx, f.Id, searchRes)
			selectedValue := selectedRangeFilters[f.Id]
			if availableValue == nil && selectedValue == nil {
				continue
			}

			respFilters = append(respFilters, &product.Filter{
				Id:   f.Id,
				Name: f.Name,
				Value: &product.Filter_RangeFilter{
					RangeFilter: &product.RangeFilter{
						AvailableValue: availableValue,
						SelectedValue:  selectedValue,
					},
				},
			})
		case *product.Filter_SwitchFilter:
			availableSwitches := buildSwitchFilterAvailableSwitches(ctx, f.Id, searchRes)
			selectedSwitch := selectedSwitchFilters[f.Id]
			if len(availableSwitches) == 0 && len(selectedSwitch) == 0 {
				continue
			}
			respFilters = append(respFilters, &product.Filter{
				Id:   f.Id,
				Name: f.Name,
				Value: &product.Filter_SwitchFilter{
					SwitchFilter: &product.SwitchFilter{
						AvailableSwitches: availableSwitches,
						SelectedSwitch:    selectedSwitch,
					},
				},
			})
		default:
			ctxzap.Extract(ctx).Warn("unknown filter type", zap.String("filter", f.Id))
			continue
		}
	}
	return respFilters
}

func buildMapSelectedFilters(ctx context.Context, filters []*product.Filter) (
	selectedListFilters map[string][]string, selectedRangeFilters map[string]*product.RangeValue, selectedSwitchFilters map[string]string) {

	selectedListFilters = map[string][]string{}
	selectedRangeFilters = map[string]*product.RangeValue{}
	selectedSwitchFilters = map[string]string{}
	for _, f := range filters {
		switch v := f.Value.(type) {
		case *product.Filter_ListFilter:
			setSelectedItems := map[string]struct{}{}
			selectedItems := make([]string, 0, len(v.ListFilter.SelectedItems))
			for _, item := range v.ListFilter.SelectedItems {
				if _, ok := setSelectedItems[item]; ok {
					continue
				}
				setSelectedItems[item] = struct{}{}
				selectedItems = append(selectedItems, item)
			}
			selectedListFilters[f.Id] = selectedItems
		case *product.Filter_RangeFilter:
			selectedRangeFilters[f.Id] = v.RangeFilter.SelectedValue
		case *product.Filter_SwitchFilter:
			selectedSwitchFilters[f.Id] = v.SwitchFilter.SelectedSwitch
		default:
			ctxzap.Extract(ctx).Warn("unknown filter type", zap.String("filter", f.Id))
			continue
		}
	}
	return
}

func extractFilterBucket(ctx context.Context, id string, searchRes *elastic.SearchResult) *elastic.AggregationBucketKeyItem {
	aggRes, ok := searchRes.Aggregations.Filters(id)
	if !ok {
		ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", id))
		return nil
	}
	if len(aggRes.Buckets) == 0 {
		return nil
	}
	return aggRes.Buckets[0]
}

func buildListFilterAvailableItems(ctx context.Context, id string, searchRes *elastic.SearchResult) []string {
	bucket := extractFilterBucket(ctx, id, searchRes)
	if bucket == nil {
		return []string{}
	}

	aggRes, ok := bucket.Aggregations.Filters(id)
	if !ok {
		ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", id))
		return []string{}
	}

	availableItems := make([]string, 0, len(aggRes.Buckets))
	for _, b := range aggRes.Buckets {
		availableItems = append(availableItems, b.Key.(string))
	}

	sort.Strings(availableItems)
	return availableItems
}

func buildRangeFilterAvailableValue(ctx context.Context, id string, searchRes *elastic.SearchResult) *product.RangeValue {
	bucket := extractFilterBucket(ctx, id, searchRes)
	if bucket == nil {
		return nil
	}

	aggResMax, ok := bucket.Aggregations.Max(makeFilterRangeMaxName(id))
	if !ok {
		ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", id))
		return nil
	}
	aggResMin, ok := bucket.Aggregations.Min(makeFilterRangeMinName(id))
	if !ok {
		ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", id))
		return nil
	}
	return &product.RangeValue{
		Max: int64(*aggResMax.Value),
		Min: int64(*aggResMin.Value),
	}
}

func buildSwitchFilterAvailableSwitches(ctx context.Context, id string, searchRes *elastic.SearchResult) []string {
	bucket := extractFilterBucket(ctx, id, searchRes)
	if bucket == nil {
		return []string{}
	}

	aggRes, ok := bucket.Aggregations.Filters(id)
	if !ok {
		ctxzap.Extract(ctx).Warn("filter not found", zap.String("filter", id))
		return []string{}
	}
	availableSwitches := make([]string, 0, len(aggRes.Buckets))
	for _, b := range aggRes.Buckets {
		availableSwitches = append(availableSwitches, dictSwitchFilter.getValue(id, b.Key))
	}

	sort.Strings(availableSwitches)
	return availableSwitches
}

func buildCategories(ctx context.Context, categoryId string, searchRes *elastic.SearchResult) []*product.Category {
	aggRes, ok := searchRes.Aggregations.Filters(fieldCategories)
	if !ok {
		ctxzap.Extract(ctx).Warn("categories not found")
		return []*product.Category{}
	}

	if len(aggRes.Buckets) == 0 {
		return []*product.Category{}
	}
	bucket := aggRes.Buckets[0]

	aggRes, ok = bucket.Aggregations.Filters(fieldCategories)
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
