package main

import (
	_ "github.com/labstack/echo"
	"go_rpc/helper"
	"go_rpc/services"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {

	//credos, err := credentials.NewServerTLSFromFile("./keys/server.crt", "./keys/server.key")
	//if err != nil {
	//	log.Fatalf("credentials.NewServerTLSFromFile:【%v】", err)
	//}

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredos()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	log.Println(rpcServer.GetServiceInfo())
	/*
			TCP协议
		//lis, err := net.Listen("tcp", ":8881")
		//if err != nil {
		//	log.Printf("net.Listen:【%v】", err)
		//}
		//err = rpcServer.Serve(lis)
		//if err != nil {
		//	log.Printf("rpcServer.Serve:【%v】", err)
		//}
	*/

	/*
		HTTP协议
	*/

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcServer.ServeHTTP(writer, request)
		log.Println(request)
	})
	httpServer := &http.Server{
		Addr:    ":8881",
		Handler: mux,
	}
	err := httpServer.ListenAndServeTLS("./cert/server.pem", "./cert/server.key")
	if err != nil {
		log.Printf("httpServer.ListenAndServeTLS:【%v】", err)
	}
}
