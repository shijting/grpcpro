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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cred, err := credentials.NewClientTLSFromFile("configs/certs/server.pem", "test.grpc.sjt.com")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.DialContext(ctx,
		":8080",
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
