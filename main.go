package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"exemploserversidetlsclient/src/pb/products"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCA, err := os.ReadFile("./src/cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCA) {
		return nil, fmt.Errorf("erro ao adicionar o certificado no pool")
	}

	clientCert, err := tls.LoadX509KeyPair("./src/cert/client-cert.pem", "./src/cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.NewClient("0.0.0.0:8080", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := products.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	productList, err := client.FindAll(ctx, &products.ListProductRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("products: %+v\n", productList)
}
