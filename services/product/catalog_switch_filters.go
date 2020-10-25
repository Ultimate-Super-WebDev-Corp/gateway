package product

import (
	"github.com/olivere/elastic/v7"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
)

func init() {
	for _, f := range dictFilters {
		switch f.Value.(type) {
		case *product.Filter_SwitchFilter:
			if _, ok := dictSwitchFilter.dict[f.Id]; !ok {
				panic("dictFilters or dictSwitchFilter inconsistent state")
			}
		}
	}
}

type switchFilterValue interface {
	getValue(value interface{}) string
	getEQuery(filterId string, sw string) elastic.Query
}

type strDictSwitchFilter struct {
	dict map[string]switchFilterValue
}

func (d strDictSwitchFilter) getValue(filterId string, value interface{}) string {
	return d.dict[filterId].getValue(value)
}

func (d strDictSwitchFilter) getEQuery(filterId string, sw string) elastic.Query {
	return d.dict[filterId].getEQuery(filterId, sw)
}

var dictSwitchFilter = strDictSwitchFilter{
	dict: map[string]switchFilterValue{
		fieldRating: gteSwitchFilterValue{
			review.Rating_UNDEFINED:            "от ☆☆☆☆☆",
			review.Rating_HALF_STARS:           "от ☆☆☆☆☆",
			review.Rating_ONE_STARS:            "от ★☆☆☆☆",
			review.Rating_ONE_AND_HALF_STARS:   "от ★☆☆☆☆",
			review.Rating_TWO_STARS:            "от ★★☆☆☆",
			review.Rating_TWO_AND_HALF_STARS:   "от ★★☆☆☆",
			review.Rating_THREE_STARS:          "от ★★★☆☆",
			review.Rating_THREE_AND_HALF_STARS: "от ★★★☆☆",
			review.Rating_FOUR_STARS:           "от ★★★★☆",
			review.Rating_FOUR_AND_HALF_STARS:  "от ★★★★☆",
			review.Rating_FIVE_STARS:           "от ★★★★★",
		},
	},
}

type gteSwitchFilterValue map[review.Rating]string

func (s gteSwitchFilterValue) getValue(value interface{}) string {
	return s[review.Rating(value.(float64))]
}

func (s gteSwitchFilterValue) getEQuery(filterId string, sw string) elastic.Query {
	key := review.Rating(1000) // invalid rating value
	for k, v := range s {
		if v == sw && key > k {
			key = k
		}
	}

	if key == 1000 {
		return nil
	}

	return elastic.NewRangeQuery(filterId).Gte(key)
}
