package main

import (
	"context"
	"exemploserversidetlsclient/src/pb/products"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {

}

func main() {
	conn, err := grpc.NewClient("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := products.NewProductServiceClient(conn)
	productList, err := client.FindAll(context.Background(), &products.ListProductRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("products: %+v\n", productList)
}
