package main

import (
	"context"
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes called")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling primes %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading from stream %v\n", err)
		}
		log.Printf("Primes: %d\n", res.Result)
	}
}
