package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := file.NewFileClient(conn)
	stream, err := c.Upload(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	f, err := os.Open("clients_test/file_uploader/photo_2020-09-22_23-05-13.jpg")
	if err != nil {
		log.Fatalf(err.Error())
	}

	ch := &file.Chunk{
		OneOfChunk: &file.Chunk_Meta{
			Meta: &file.FileMetadata{
				Type: file.FileType_JPEG,
			},
		},
	}
	err = stream.Send(ch)

	if err != nil {
		log.Fatalf(err.Error())
	}

	res, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf(err.Error())
	}

	l := len(res) / 2
	ch = &file.Chunk{
		OneOfChunk: &file.Chunk_Chunk{
			Chunk: res[:l],
		},
	}

	err = stream.Send(ch)

	ch = &file.Chunk{
		OneOfChunk: &file.Chunk_Chunk{
			Chunk: res[l:],
		},
	}

	err = stream.Send(ch)

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(reply.UUID)
}
