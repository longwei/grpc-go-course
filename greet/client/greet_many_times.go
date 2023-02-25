package main

import (
	"context"
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes called")

	req := &pb.GreetRequest{
		FirstName: "longwei",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("Could not GreetManyTimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("error while reading from stream %v\n", err)
		}
		log.Printf("GreetManyTime: %s\n", msg.Result)
	}
}
