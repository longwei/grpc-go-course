package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet funcntion was invoked with %v", in)
	return &pb.GreetResponse{
		Result: "hello from server" + in.FirstName,
	}, nil
}
