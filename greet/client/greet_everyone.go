package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone called")

	stream, err := c.GreetEveryone(context.Background())
	log.Println("get the stream...")

	if err != nil {
		log.Fatal("Could not doGreetEveryone %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "longwei1"},
		{FirstName: "longwei2"},
		{FirstName: "longwei3"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending requests: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
