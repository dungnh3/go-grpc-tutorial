package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (s *server) SumWithDealine(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Print("sum with deadline.. \n")
	result := in.Number1 + in.Number2
	time.Sleep(3 * time.Second)
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

func (s *server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("average call..")
	var total int32
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var result float32
			result = float32(total) / float32(count)
			resp := &pb.AverageResponse{
				Result: result,
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatalf("error while average %v", err)
		}
		total = total + req.Number
		count++
	}
	return nil
}

func (s *server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("find max call..")
	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		num := req.Number
		log.Printf("recieve number from clien %v \n", num)
		if num > max {
			max = num
		}
		err = stream.Send(&pb.FindMaxResponse{
			Result: max,
		})
		if err != nil {
			return err
		}
	}
}

func (s *server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Print("sum sqrt \n")
	num := in.Number
	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "input number must great than 0")
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
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
