package main

import (
	"go_rpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {

	creds, err := credentials.NewServerTLSFromFile("./keys/server.crt", "./keys/server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile:【%v】", err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	log.Println(rpcServer.GetServiceInfo())

	lis, err := net.Listen("tcp", ":8881")
	if err != nil {
		log.Printf("net.Listen:【%v】", err)
	}
	err = rpcServer.Serve(lis)
	if err != nil {
		log.Printf("rpcServer.Serve:【%v】", err)
	}
}
