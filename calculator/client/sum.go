package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum called")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatal("Could not sum %v\n", err)
	}

	log.Printf("response from server: Sum = %d\n", res.Result)
}
