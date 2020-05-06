package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"log"
	"net"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	resp := &greetpb.GreetResponse{
		Result: "hello " + firstName,
	}
	return resp, nil
}

func main() {
	// create a listener
	//50051 is the default port for gRPC
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	//register a service
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("valied in serve the listener %s", err.Error())
	}

}
