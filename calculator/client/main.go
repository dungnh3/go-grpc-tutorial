package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	clientConn, err := grpc.Dial("localhost:10069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer clientConn.Close()

	client := pb.NewCalculatorServiceClient(clientConn)
	//callSum(client)
	callPrimeNumberDecomposition(client)

	log.Println("service client %v", client)
}

func callSum(client pb.CalculatorServiceClient) {
	resp, err := client.Sum(context.Background(), &pb.SumRequest{
		Number1: 5,
		Number2: 6,
	})

	if err != nil {
		log.Fatalf("call sum api error %v", err)
	}

	log.Printf("sum api response %v \n", resp.Result)
}

func callPrimeNumberDecomposition(client pb.CalculatorServiceClient) {
	stream, err := client.PrimeNumberDecomposition(context.Background(), &pb.PNDRequest{Number: 120})
	if err != nil {
		log.Fatalf("call pnd api error %v", err)
	}
	if err != nil {
		log.Fatalf("call PND failed %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("server finish streaming")
			return
		}
		log.Println("prime number response %v", resp.Result)
	}
}
