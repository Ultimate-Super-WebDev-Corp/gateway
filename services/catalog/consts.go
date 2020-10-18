package catalog

import "github.com/olivere/elastic/v7"

const (
	objectProduct = "product"

	fieldId    = "id"
	fieldBrand = "brand"

	elasticFilterFieldBrand = "brand.keyword"
)

type filterList struct {
	name        string
	aggregation elastic.Aggregation
}

var filtersList = []filterList{
	{
		name:        fieldBrand,
		aggregation: elastic.NewTermsAggregation().Field(getElasticFilterField(fieldBrand)),
	},
}

var fieldToElasticFilterField = map[string]string{
	fieldBrand: elasticFilterFieldBrand,
}

func getElasticFilterField(field string) string {
	elasticField, ok := fieldToElasticFilterField[field]
	if ok {
		return elasticField
	}
	return field
}
