package main

import (
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"io"
	"log"
	"net"
)

type server struct {
}

func (s *server) Greet(req greetpb.GreetClientStream_GreetServer) error {
	result := ""
	for {
		message, err := req.Recv()
		if err == io.EOF {
			err2 := req.SendAndClose(&greetpb.GreetClientResponse{
				Result: result,
			})
			if err2 != nil {
				log.Fatal(err)
			}
			break

		}
		if err != nil {
			log.Fatal(err)
		}
		result += "hello " + message.Greeting.GetName() + "! "
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetClientStreamServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
