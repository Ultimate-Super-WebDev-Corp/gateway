package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

func TestProductGetByID(t *testing.T) {
	sessionWithCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21lcklkIjoxMDAwLCJQYXNzd29yZElkIjoyLCJVcGRhdGVkQXQiOjI2MDEzMDc3MzIzNTQzOTE2MDB9.P1fryf6q1kmwZMuF8vEkAaRBQuoj7uRxuTZLZD7GxiE"
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithCusId))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	productCli := product.NewProductClient(conn)
	respGetByID, err := productCli.GetByID(ctx, &product.GetByIDRequest{
		Id: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, respGetByID.Id)
}
