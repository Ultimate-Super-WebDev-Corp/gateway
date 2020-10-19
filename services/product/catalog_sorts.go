package product

import "github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"

var dictSorts = []*product.Sort{
	{
		Field: fieldRating,
		Name:  "по рейтингу",
	},
	{
		Field: fieldVotes,
		Name:  "по количеству отзывов",
	},
}
