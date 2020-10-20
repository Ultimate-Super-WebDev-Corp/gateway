package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

const (
	address = "localhost:8080"
)

func main() {
	//fmt.Println(fmt.Println(time.Now().UTC().UnixNano()))
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//defer conn.Close()
	//
	//c := product.NewProductClient(conn)
	//
	//ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
	//	"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IiIsIlVwZGF0ZWRBdCI6MjYwMTMwNzczMjM1NDM5MTYwMH0.9X89JDfmp1pfG-j2nTEx67C04ojg2xyi1b3GAK9haYs"},
	//})
	//
	//var header metadata.MD
	//var header2 metadata.MD
	//
	//resp, _ := c.SearchByUUIDs(ctx, &product.SearchByUUIDsRequest{
	//	UUIDs: []string{"bf5b6d56-0262-11eb-8ebb-c0b6f983d777.20200929"},
	//}, grpc.Header(&header), grpc.Header(&header2))
	//
	//_ = resp
	//_ = header
	//_ = header2
	//fmt.Println("wtf")
	productInsert()
}

type aDermaPrd struct {
	Images      []string `json:"images"`
	Name        string   `json:"eng_name"`
	Description string   `json:"description"`
	Country     string   `json:"country"`
	Brand       string   `json:"brand"`
}

var cat = [][]string{
	{"face"},
	{"hair"},
	{"face", "mask"},
	{"hair", "hair_conditioner"},
	{"face", "mask"},
	{"hair", "shampoo"},
	{"face", "cream"},
}

func productInsert() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := product.NewProductClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"token": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJQYXNzd29yZElkIjotMSwiQ3VzdG9tZXJJZCI6MSwiVXBkYXRlZEF0IjoyNjAxMzA3NzMyMzU0MzkxNjAwfQ.AZsSe6JsFD_BAIHuVzCksj00wJJo3rfhJrkmSqjhJjo"},
	})

	f, err := os.Open("clients_test/product/1005767.json")
	_ = err
	allF, err := ioutil.ReadAll(f)
	_ = err

	aDerms := make([][]aDermaPrd, 0, 1000)

	err = json.Unmarshal(allF, &aDerms)
	fmt.Println(err)
	for _, derm := range aDerms {
		for i, d := range derm {
			_, err := c.Create(ctx, &product.ProductMsg{
				Name:        d.Name,
				Brand:       d.Brand,
				Description: d.Description,
				Images:      d.Images,
				Country:     d.Country,
				CategoryIds: cat[i%(len(cat))],
			})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
