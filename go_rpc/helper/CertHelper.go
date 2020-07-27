package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func GetServerCredos() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("./cert/server.pem", "./cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
}

func GetClientCredos() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("./cert/client.pem", "./cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //加载客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
}
