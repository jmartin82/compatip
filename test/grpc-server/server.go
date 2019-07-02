package main

import (
	"context"
	"net"
	"os"

	"github.com/jmartin82/compatip/test/grpc-server/rpc"
	"google.golang.org/grpc"
)

type VersionService struct{}

func (us *VersionService) Current(context.Context, *rpc.Empty) (*rpc.VersionMessage, error) {
	return &rpc.VersionMessage{
		Version: "1.4.23",
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		os.Exit(1)
	}
	server := grpc.NewServer()
	service := &VersionService{}
	rpc.RegisterVersionServer(server, service)

	err = server.Serve(l)
	if err != nil {
		os.Exit(1)
	}
}
