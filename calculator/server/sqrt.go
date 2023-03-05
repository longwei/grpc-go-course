package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function called with %v", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received a negative number: %d", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
