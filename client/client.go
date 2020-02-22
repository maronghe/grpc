package main

import (
	"fmt"
	"grpc/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Username: "Logan Ma"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.Message)

	rb, err := c.SayBye(context.Background(), &pb.HelloRequest{Username: "Logan Ma"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rb.Message)
}
