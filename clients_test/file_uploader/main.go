package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IiJ9.IJ9LcRwoX0nDaPKNgI-SvbmqIrLjjc0_rfp40Of65_k"},
	})

	header := metadata.MD{}
	stream, err := c.Upload(ctx, grpc.Header(&header))
	if err != nil {
		log.Fatalf(err.Error())
	}

	f, err := os.Open("clients_test/file_uploader/photo_2020-09-21_22-44-27.jpg")
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
