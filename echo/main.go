package main

import (
	"fmt"
	"log"
	"net"

	"github.com/troydai/blocks/echo/proto"
	"github.com/troydai/blocks/echo/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5436")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	s, err := server.New()
	if err != nil {
		log.Fatalf("fail to create server: %v", err)
	}
	proto.RegisterEchoServerServer(grpcServer, s)

	fmt.Println("server starting ...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to start server: %v", err)
	}
}
