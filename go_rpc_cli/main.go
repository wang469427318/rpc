package main

import (
	"context"
	"go_rpc_cli/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("./keys/server.crt", "andy")

	conn, err := grpc.Dial(":8881", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("grpc.Dial【%v】", err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatalf("prodClient.GetProdStock【%v】", err)
	}
	log.Println(prodRes.ProdStock)
}
