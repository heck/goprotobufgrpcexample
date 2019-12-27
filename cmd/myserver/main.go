package main

// stolen from: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

import (
	"context"
	"log"
	"net"

	pb "github.com/heck/goprotobufgrpcexample/api/mypersonpkg"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedMyPersonServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPerson(ctx context.Context, in *pb.MyPersonRequest) (*pb.MyPersonResponse, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.MyPersonResponse{
		Id:   in.GetId(),
		Name: "Harvey Mud",
		Age:  107,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyPersonServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
