package main

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/product"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	grpcConn, err := srv.GrpcDial()
	if err != nil {
		panic(err)
	}

	fileCli, err := file.NewFile(file.Dependences{
		Registrar: srv.RpcServer,
		GrpcConn:  grpcConn,
	})
	if err != nil {
		panic(err)
	}

	err = product.NewProduct(product.Dependences{
		Registrar: srv.RpcServer,
		FileCli:   fileCli,
	})
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
