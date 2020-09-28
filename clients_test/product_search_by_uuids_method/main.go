package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := product.NewProductClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IiJ9.IJ9LcRwoX0nDaPKNgI-SvbmqIrLjjc0_rfp40Of65_k"},
	})

	var header metadata.MD
	var header2 metadata.MD

	_, _ = c.SearchByUUIDs(ctx, &product.SearchByUUIDsRequest{
		UUIDs: []string{"3bee0ec0-0176-11eb-b2f0-c0b6f983d777.20200928"},
	}, grpc.Header(&header), grpc.Header(&header2))

	_ = header
	_ = header2
	fmt.Println("wtf")
}
