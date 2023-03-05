package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax called")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatal("Error while opening stream: %v\n", err)
	}

	waitc := make(chan chan struct{})

	// one channel to send
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			log.Printf("sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// one channel to listen

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}
			log.Printf("Received a new maximum of...: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
