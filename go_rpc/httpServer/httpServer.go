package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go_rpc/helper"
	"go_rpc/services"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredos())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8881", opt)
	if err != nil {
		log.Fatalf("services.RegisterProdServiceHandlerFromEndpoint【%v】", err)
	}
	httpServer := &http.Server{
		Addr:    ":8880",
		Handler: gwmux,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("httpServer.ListenAndServe【%v】", err)
	}
}
