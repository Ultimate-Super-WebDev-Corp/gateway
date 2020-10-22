package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

func TestCatalogFlow(t *testing.T) {
	sessionWithCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21lcklkIjoxMDAwLCJQYXNzd29yZElkIjoyLCJVcGRhdGVkQXQiOjI2MDEzMDc3MzIzNTQzOTE2MDB9.P1fryf6q1kmwZMuF8vEkAaRBQuoj7uRxuTZLZD7GxiE"
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithCusId))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()
	productCli := product.NewProductClient(conn)

	searchResp, err := productCli.Catalog(ctx, &product.CatalogRequest{
		Filters: []*product.Filter{
			//{
			//	Id: "brand",
			//	Value: &product.Filter_ListFilter{
			//		ListFilter: &product.ListFilter{
			//			List: []string{"MAKEUP"},
			//		},
			//	},
			//},
			//{
			//	Id: "rating",
			//	Value: &product.Filter_RangeFilter{RangeFilter: &product.RangeFilter{
			//		Min: 4,
			//		Max: 6,
			//	}},
			//},
			//{
			//	Id: "rating",
			//	Value: &product.Filter_SwitchFilter{
			//		SwitchFilter: &product.SwitchFilter{
			//			Switches: []string{"от ★★★★☆"},
			//		},
			//	},
			//},
		},
		//TextSearch: "Auriga",
		Token: 0,
		Limit: 4,
	})
	_, _ = searchResp, err

	fmt.Println(searchResp)
}
