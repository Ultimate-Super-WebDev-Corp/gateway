package main

import (
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/file_uploader"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/services/search"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	if err := search.NewSearch(search.Dependences{
		Registrar: srv.RpcServer,
	}); err != nil {
		panic(err)
	}

	if err := file_uploader.NewFileUploader(file_uploader.Dependences{
		Registrar: srv.RpcServer,
	}); err != nil {
		panic(err)
	}

	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
