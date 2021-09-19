package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/troydai/blocks/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5436", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial tcp: %v", err)
	}
	defer conn.Close()

	client := proto.NewEchoServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &proto.EchoRequest{Message: "well, i'm waiting."})
	if err != nil {
		log.Fatalf("fail to communicate with the server")
	}
	fmt.Printf("message: %s\n", resp.Message)
	fmt.Printf(" digest: %s\n", resp.Digest)
}
