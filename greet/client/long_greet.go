package main

import (
	"context"
	"log"
	"time"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet called")

	reqs := []*pb.GreetRequest{
		{FirstName: "longwei1"},
		{FirstName: "longwei2"},
		{FirstName: "longwei3"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Could not LongGreet Times %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while receing response from LongGreet %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
