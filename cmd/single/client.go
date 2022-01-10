package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/shijting/grpcpro/pgfiles"

	"google.golang.org/grpc"
)

const (
	certFile           = "configs/certs/server.pem"
	serverNameOverride = "test.grpc.sjt.com"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cred, err := credentials.NewClientTLSFromFile(certFile, serverNameOverride)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.DialContext(ctx,
		"test.grpc.sjt.com:8080",
		grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}

	c := pgfiles.NewProdServerClient(conn)
	resp, err := c.GetProd(ctx, &pgfiles.ProdRequest{Id: 123456})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
