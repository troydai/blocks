package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/troydai/blocks/echo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	m := metadata.New(map[string]string{"reqctx-persist": "from-edge"})
	ctx = metadata.NewOutgoingContext(ctx, m)

	req := &proto.EchoRequest{Message: extractMessage()}

	resp, err := client.Echo(ctx, req)
	if err != nil {
		log.Fatalf("fail to communicate with the server")
	}
	fmt.Printf("message: %s\n", resp.Message)
	fmt.Printf(" digest: %x\n", resp.Digest)
}

func extractMessage() string {
	if len(os.Args) < 2 {
		return "test message"
	}
	return os.Args[1]
}
