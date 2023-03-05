package main

import (
	"context"
	"log"
	"time"

	pb "github.com/longwei/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadLine called")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "longwei",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		log.Fatal("Could not doGreetWithDeadline %v\n", err)
		e, ok := status.FromError(err)
		if ok {
			// grpc
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Println("Unexpected gRPC error %v\n", err)
			}

		} else {
			//non grpc
			log.Fatalf("a non grpc error: %v\n", err)
		}

	}

	log.Printf("doGreetWithDeadline from server: %s\n", res.Result)
}
