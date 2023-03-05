package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, in int32) {
	log.Println("doSqrt called")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: in,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("A grpc error message %s\n", e.Message())
			log.Printf("A grpc error code %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("oops, sent negative number")
				return
			}
		} else {
			log.Fatalf("A non grpc error %v\n", err)
		}
	}

	log.Printf("response from server: sqrt = %f\n", res.Result)
}
