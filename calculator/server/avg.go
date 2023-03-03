package main

import (
	"io"
	"log"

	pb "github.com/longwei/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was called\n")

	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client %v\n", err)
		}
		log.Printf("Receiving number: %d\n", req.Number)
		sum += req.Number
		count += 1
	}
}

// func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
// 	log.Printf("primes function was called %v\n", in)

// 	number := in.Number
// 	divisor := int64(2)
// 	for number > 1 {
// 		if number%divisor == 0 {
// 			stream.Send(&pb.PrimeResponse{
// 				Result: divisor,
// 			})
// 			number /= divisor
// 		} else {
// 			divisor++
// 		}
// 	}
// 	return nil
// }
