package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

const (
	address = "localhost:8080"
)

func main() {
	fmt.Println(fmt.Println(time.Now().UTC().UnixNano()))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := product.NewProductClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IiIsIlVwZGF0ZWRBdCI6MTYwMTMwNzczMjM1NDM5MTMwMH0.7Ny4fkVbT1N55E7sziE_Lh957CdltNroEnLW6_4kZek"},
	})

	var header metadata.MD
	var header2 metadata.MD

	_, _ = c.SearchByUUIDs(ctx, &product.SearchByUUIDsRequest{
		UUIDs: []string{"0433ba0e-019f-11eb-a2c4-c0b6f983d777.20200928"},
	}, grpc.Header(&header), grpc.Header(&header2))

	_ = header
	_ = header2
	fmt.Println("wtf")
}
