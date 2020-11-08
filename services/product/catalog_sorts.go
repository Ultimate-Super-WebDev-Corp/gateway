package product

import "github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"

var defaultSort = &product.Sort{
	Id:        fieldRating,
	Name:      "По рейтингу",
	Ascending: false,
}

var dictSorts = []*product.Sort{
	{
		Id:        fieldRating,
		Name:      "По рейтингу",
		Ascending: false,
	},
	{
		Id:        fieldVotes,
		Name:      "По количеству отзывов",
		Ascending: false,
	},
}
