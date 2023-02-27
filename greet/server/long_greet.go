package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet called")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("error while reading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)

	}

}
