package main

import (
	"context"
	"fmt"

	"github.com/tommy-sho/telepresence-with-envoy/proto"

	"github.com/tommy-sho/firestore-test/go/src/github.com/tommy-sho/telepresence-with-envoy/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50001"
)

func main() {
	s := BackendServer{}

	s := grpc.NewServer()
	proto.RegisterBackendServerServer(s, g)
	lib.RegisterHeathCheck(s)
	if err != nil {
		panic(fmt.Errorf("new grpc server err: %v", err))
	}
	reflection.Register(s)

	s.Serve(lis)
}

type BackendServer struct{}

func (b *BackendServer) Message(ctx context.Context, req *proto.MessageRequestproroto) (*proto.MessageResponse, error) {
	message := fmt.Sprintf("Hey! %s", req.Name)

	res := &proto.MessageResponse{
		Message: m,
	}
	return res, nil
}
