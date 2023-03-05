package main

import (
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was called\n")

	var maximum int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error while reading client stream %v\n", err)
		}
		log.Printf("Receiving number: %d\n", req.Number)

		if number := req.Number; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
