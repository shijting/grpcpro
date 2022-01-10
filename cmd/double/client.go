package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/shijting/grpcpro/pgfiles"

	"google.golang.org/grpc"
)

const (
	serverNameOverride = "test.grpc.sjt.com"
	serverPort         = 8080
	clientCertFile     = "configs/certs/client.crt"
	clientKeyFile      = "configs/certs/client.key"
	caRootFile         = "configs/certs/ca.crt"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cret, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatalln(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caRootFile)
	if err != nil {
		log.Fatalln(err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret}, // 放入客户端证书
		ServerName:   serverNameOverride,
		RootCAs:      certPool,
	})

	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf("%s:%d", serverNameOverride, serverPort),
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	c := pgfiles.NewProdServerClient(conn)
	resp, err := c.GetProd(ctx, &pgfiles.ProdRequest{Id: 200})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
