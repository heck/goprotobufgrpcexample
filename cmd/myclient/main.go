package main

// stolen from: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go

import (
	"context"
	"log"
	"time"

	pb "github.com/heck/goprotobufgrpcexample/api/mypersonpkg"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyPersonServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPerson(ctx, &pb.MyPersonRequest{Id: 123})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Person: %#v\n", r)
}
