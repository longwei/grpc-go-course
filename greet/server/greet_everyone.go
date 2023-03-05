package main

import (
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("Error while reading from client stream %v\n", err)
		}
		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		log.Printf("Sending: %v\n", res)
		if err != nil {
			log.Fatal("Error while sending to client %v\n", err)
		}

	}
}
