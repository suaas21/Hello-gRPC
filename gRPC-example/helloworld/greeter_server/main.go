package main

import (
	"context"
	pb "github.com/Hello-gRPC/gRPC-example/helloworld/helloworld"
    "google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error)  {
	log.Printf("Received: %v", in.GetName())
	return &pb.Response{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server %v", err)
	}
}
