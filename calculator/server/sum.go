package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function called with %v", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
