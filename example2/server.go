package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "envoy_example/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
    fmt.Printf("Server received Hello from %v %v...\n", ctx, request)
	response := &pb.HelloResponse{Greeting: "Hello " + name}
	return response, nil
}

func main() {
	address := "0.0.0.0:5050"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Backend Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...\n", address)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
