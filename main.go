package main

import (
	"go.uber.org/zap"

	"github.com/Ultimate-Super-WebDev-Corp/server/server"
	"github.com/Ultimate-Super-WebDev-Corp/server/services/customer"
)

func main() {
	srv := server.NewServer()

	err :=customer.NewCustomer(customer.Dependences{
		Registrar: srv.RpcServer,
	})
	if err!= nil{
		srv.Logger.Panic("create service customer error", zap.Any("error", err))
	}

	if err := srv.Serve(); err != nil {
		srv.Logger.Panic("failed to serve", zap.Any("error", err))
	}
}
