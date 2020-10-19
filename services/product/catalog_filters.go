package product

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

var dictFilters = []product.Filter{
	{
		Field: fieldBrand,
		Name:  "Бренд",
		Value: &product.Filter_ListFilter{},
	},
	{
		Field: fieldRating,
		Name:  "Рейтинг",
		Value: &product.Filter_RangeFilter{},
	},
	{
		Field: fieldVotes,
		Name:  "Количество отзывов",
		Value: &product.Filter_RangeFilter{},
	},
}

var fieldToEFilterField = map[string]string{
	fieldBrand: eFilterFieldBrand,
}

func getEFilterField(field string) string {
	elasticField, ok := fieldToEFilterField[field]
	if ok {
		return elasticField
	}
	return field
}

func makeFilterRangeMinName(filterName string) string {
	return "min_" + filterName
}

func makeFilterRangeMaxName(filterName string) string {
	return "max_" + filterName
}
