package product

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

var dictFilters = []product.Filter{
	{
		Id:    fieldBrand,
		Name:  "Бренд",
		Value: &product.Filter_ListFilter{},
	},
	{
		Id:    fieldRating,
		Name:  "Рейтинг",
		Value: &product.Filter_RangeFilter{},
	},
	{
		Id:    fieldVotes,
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
