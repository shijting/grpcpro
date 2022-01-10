package main

import (
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/shijting/grpcpro/pgfiles"
	"github.com/shijting/grpcpro/pkg/prod"
	"google.golang.org/grpc"
)

const (
	serverKeyFile  = "configs/certs/server.key"
	serverCertFile = "configs/certs/server.pem"
)

func main() {
	cred, err := credentials.NewServerTLSFromFile(serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer(grpc.Creds(cred))
	pgfiles.RegisterProdServerServer(srv, prod.NewProdServer())
	lis, _ := net.Listen("tcp", ":8080")
	if err := srv.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
