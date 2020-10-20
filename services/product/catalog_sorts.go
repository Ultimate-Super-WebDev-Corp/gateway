package product

import "github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"

var defaultSort = &product.Sort{
	Id:        fieldRating,
	Name:      "по рейтингу",
	Ascending: false,
}

var dictSorts = []*product.Sort{
	{
		Id:   fieldRating,
		Name: "по рейтингу",
	},
	{
		Id:   fieldVotes,
		Name: "по количеству отзывов",
	},
}
