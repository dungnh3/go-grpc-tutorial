package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	clientConn, err := grpc.Dial("localhost:10069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer clientConn.Close()

	client := pb.NewCalculatorServiceClient(clientConn)
	callSum(client)

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
