package main

import (
	"context"
	pb "github.com/dungnh3/go-grpc-tutorial/grpc_gateway/demopb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Echo(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{
		Msg: func() string {
			var i int32
			for i = 0; i < in.Number; i++ {
				in.Msg = in.Msg + " hello"
			}
			return in.Msg
		}(),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:10443")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_validator.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_validator.StreamServerInterceptor(),
		),
	)

	pb.RegisterMessageServiceServer(s, &server{})

	log.Println("grpc gateway is running..")
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
