package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	clientConn, err := grpc.Dial("localhost:10069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer clientConn.Close()

	client := pb.NewCalculatorServiceClient(clientConn)

	// unary => open comment before callSum function below if you want to test
	//callSum(client)

	// streaming server => open comment before callPrimeNumberDecomposition function below if you want to test
	//callPrimeNumberDecomposition(client)

	// streaming client => open comment before callAverage function below if you want to test
	//callAverage(client)

	callMax(client)

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

func callAverage(client pb.CalculatorServiceClient) {
	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalf("call average api error %v", err)
	}
	arr := []int32{1, 2, 3, 4, 5, 6}
	for _, num := range arr {
		err := stream.Send(&pb.AverageRequest{
			Number: num,
		})
		if err != nil {
			log.Fatalf("send number to server fail %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("recieve result from server fail %v", err)
	}
	log.Printf("result average %v", res.Result)
}

func callMax(client pb.CalculatorServiceClient) {
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("call average api error %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(stream pb.CalculatorService_MaxClient, wg *sync.WaitGroup) {
		defer wg.Done()
		arr := []int32{1, 2, 3, 4, 5, 6, 100, 50, 200, 1000, 500, 400, 2000}
		for _, num := range arr {
			err := stream.Send(&pb.FindMaxRequest{
				Number: num,
			})
			if err != nil {
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}(stream, &wg)

	go func(stream pb.CalculatorService_MaxClient, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("recieve max from server fail %v", err)
				break
			}
			log.Printf("result max %v", resp.Result)
		}
	}(stream, &wg)

	wg.Wait()
}
