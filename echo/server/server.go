package server

import (
	"context"

	"github.com/troydai/blocks/echo/proto"
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
	return &proto.EchoResponse{
		Message: req.Message,
		Digest:  "simple echo server",
	}, nil
}
