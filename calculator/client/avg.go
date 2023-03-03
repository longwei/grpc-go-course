package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg called")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatal("Could not sum %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("error while receiving response %v\n", err)
	}

	log.Printf("Avg response from server: Sum = %f\n", res.Result)
}
