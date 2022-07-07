package main

import (
	"context"
	"fmt"
	"log"
	
	pb "envoy_example/protos"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("host.docker.internal:1337", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewHelloServiceClient(cc)
	request := &pb.HelloRequest{Name: "brian"}

	ctx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Bearer foo", "Bar", "baz")

	resp, err := client.Hello(ctx, request)
	if err != nil {
		errStatus, isGrpcErr := status.FromError(err)
		if !isGrpcErr {
			fmt.Printf("Unknown error! %v\n", errStatus.Message())
			return
		}
		code := errStatus.Code()
		msg := errStatus.Message()
		fmt.Println(code)
		fmt.Println(msg)
	} else {
		fmt.Printf("Receive response => [%v]\n", resp.Greeting)
	}
}
