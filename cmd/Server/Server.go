package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "simple/api"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSimpleServicesServer
}

func (s *server) GetDeviceInterfaces(ctx context.Context, in *pb.RequestR) (*pb.ResponseR, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ResponseR{Name: "McCarthy " + in.GetName()}, nil
}

func main() {
	fmt.Println("hello Server")
	fmt.Println("starting server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("err:", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimpleServicesServer(s, &server{})

	fmt.Println("About to run serve")

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ")
	}

}
