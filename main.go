package main

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/customer"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	err = customer.NewCustomer(customer.Dependences{
		Registrar: srv.RpcServer,
	})
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
