package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/grpc_gateway/demopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Echo(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{
		Msg: in.Msg,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:10443")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMessageServiceServer(s, &server{})

	log.Println("grpc gateway is running..")
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
