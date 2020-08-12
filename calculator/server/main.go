package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Print("sum call \n")
	result := in.Number1 + in.Number2
	return &pb.SumResponse{
		Result: result,
	}, nil
}

func (s *server) PrimeNumberDecomposition(req *pb.PNDRequest,
	stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("prime number decomposition call..")
	num := req.Number
	k := int32(2)
	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PNDResponse{
				Result: k,
			})
			num = num / k
			time.Sleep(1 * time.Second)
		} else {
			k++
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:10069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &server{})

	log.Println("calculator is running..")
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
