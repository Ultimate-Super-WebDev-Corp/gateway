package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/catalog"
)

func TestCatalogFlow(t *testing.T) {
	sessionWithCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21lcklkIjoxMDAwLCJQYXNzd29yZElkIjoyLCJVcGRhdGVkQXQiOjI2MDEzMDc3MzIzNTQzOTE2MDB9.P1fryf6q1kmwZMuF8vEkAaRBQuoj7uRxuTZLZD7GxiE"
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithCusId))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()
	catalogCli := catalog.NewCatalogClient(conn)

	searchResp, err := catalogCli.Search(ctx, &catalog.SearchRequest{
		Sort: &catalog.Sort{
			OrderBy:   0,
			Ascending: false,
		},
		Filter: []*catalog.Filter{
			{
				Field: "brand",
				List:  []string{"Biolage"},
			},
		},
		Token: 0,
		Limit: 4,
	})
	_, _ = searchResp, err

	fmt.Println(searchResp)
}
