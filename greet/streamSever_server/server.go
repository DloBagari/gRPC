package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"log"
	"net"
	"time"
)

type server struct {
}

func (s *server) Greet(req *greetpb.GreetManyRequest, stream greetpb.GreetMany_GreetServer) error {
	firstName := req.Greeting.GetFirstName()
	for i := 0; i < 10; i++ {
		response := &greetpb.GreetManyResponse{
			Result: fmt.Sprintf("hello %s . respnose %d", firstName, i),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetManyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
