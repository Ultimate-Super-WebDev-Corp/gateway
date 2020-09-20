package main

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/search"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	fileSrv, err := file.NewFile(file.Dependences{
		Registrar: srv.RpcServer,
	})
	if err != nil {
		panic(err)
	}

	err = search.NewSearch(search.Dependences{
		Registrar: srv.RpcServer,
		FileSrv:   fileSrv,
	})
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
