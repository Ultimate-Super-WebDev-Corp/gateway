package product

import "github.com/olivere/elastic/v7"

type filterList struct {
	name        string
	aggregation elastic.Aggregation
}

var filtersList = []filterList{
	{
		name:        fieldBrand,
		aggregation: elastic.NewTermsAggregation().Field(getEFilterField(fieldBrand)),
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
