package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	pb "simple/api"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("hello Client")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServicesClient(conn)

	name := "caleb"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r, err := c.GetDeviceInterfaces(ctx, &pb.RequestR{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())

	stream, err := c.GetStreaming(ctx, &pb.RequestR{Name: name})
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
	}
	for {
		name, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		log.Println(name)
	}
}
