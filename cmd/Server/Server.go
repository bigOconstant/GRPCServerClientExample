package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "simple/api"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSimpleServicesServer
	nameList []*pb.ResponseNames
}

func (s *server) GetDeviceInterfaces(ctx context.Context, in *pb.RequestR) (*pb.ResponseR, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ResponseR{Name: "McCarthy " + in.GetName()}, nil
}

func (s *server) InitNameList() {
	s.nameList = make([]*pb.ResponseNames, 3)
	for i := 0; i < 3; i++ {
		s.nameList[i] = &pb.ResponseNames{Name: ""}
	}
	s.nameList[0].Name = "gandolf"
	s.nameList[1].Name = "the"
	s.nameList[2].Name = "grey"
}
func (s *server) GetStreaming(r *pb.RequestR, stream pb.SimpleServices_GetStreamingServer) error {
	for _, name := range s.nameList {
		fmt.Println("Sending name:", name)
		time.Sleep(1 * time.Second)
		if err := stream.Send(name); err != nil {

			return err
		}
	}
	return nil
}

func main() {
	fmt.Println("hello Server")
	fmt.Println("starting server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("err:", err)
	}
	s := grpc.NewServer()
	ss := server{}
	ss.InitNameList()

	pb.RegisterSimpleServicesServer(s, &ss)

	fmt.Println("About to run serve")

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ")
	}

}
