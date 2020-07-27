package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go_rpc_cli/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {
	//cred, err := credentials.NewClientTLSFromFile("./keys/server.crt", "andy")

	cert, _ := tls.LoadX509KeyPair("./cert/client.pem", "./cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	credos := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //加载客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8881", grpc.WithTransportCredentials(credos))
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
