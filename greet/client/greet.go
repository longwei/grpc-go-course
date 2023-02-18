package main

import (
	"context"
	"log"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet called")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "longwei",
	})

	if err != nil {
		log.Fatal("Could not greet %v\n", err)
	}

	log.Printf("Greeting echo from server: %s\n", res.Result)
}
