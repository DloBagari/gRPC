package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/greet/greetpb"
	"io"
	"log"
	"net"
)

type server struct {
}

func (s *server) Bidi(req greetpb.BidiConv_BidiServer) error {
	for {
		message, err := req.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		name := message.Person.GetName()
		if name == "" {
			return status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid user name: %s", name))
		}
		if err := req.Send(&greetpb.BidiResponse{
			Result: "hello " + message.Person.GetName()}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterBidiConvServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
