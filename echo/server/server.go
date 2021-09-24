package server

import (
	"context"
	"crypto/sha256"
	"log"

	"github.com/troydai/blocks/echo/proto"
	"google.golang.org/grpc/metadata"
)

type (
	impl struct {
		proto.UnimplementedEchoServerServer
	}
)

func New() (proto.EchoServerServer, error) {
	return &impl{}, nil
}

func (s *impl) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	log.Printf("recieve request\n")
	if m, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("\tmetadata: len = %d", m.Len())
		for k, v := range m {
			log.Printf("\t\t %v: %v", k, v)
		}
	} else {
		log.Println("\tmetadata: nil")
	}

	h := sha256.New()
	h.Write([]byte(req.Message))

	return &proto.EchoResponse{
		Message: req.Message,
		Digest:  h.Sum(nil),
	}, nil
}
