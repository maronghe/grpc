package main

import (
	"fmt"
	"net"

	"grpc/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type HelloService struct {
}

func (hs HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("你好，%s", in.Username)}, nil
}

func (hs HelloService) SayBye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("context", ctx, "in", in)
	return &pb.HelloResponse{Message: fmt.Sprintf("Bye，%s", in.Username)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, HelloService{})
	s.Serve(lis)
}
