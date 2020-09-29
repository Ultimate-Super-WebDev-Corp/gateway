package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
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

	c := customer.NewCustomerClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjMyYWNjMWQ1LTAyNWEtMTFlYi05NTVhLWMwYjZmOTgzZDc3NyIsIkN1c3RvbWVySWQiOjAsIlVwZGF0ZWRBdCI6MjYwMTM4NzIzNjM3OTY3ODcwMH0.Qh1Lt-cSMSk-acFj0Kah9bxt58yT3wequ2EsjWg6_LY"},
	})

	var header metadata.MD
	var header2 metadata.MD

	resp, _ := c.Login(ctx, &customer.LoginRequest{
		Email:    "test@mail.com",
		Password: "12345",
	}, grpc.Header(&header), grpc.Header(&header2))
	_ = resp
	_ = header
	_ = header2
	fmt.Println("wtf")
}
