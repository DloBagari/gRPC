package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := greetpb.NewGreetManyClient(conn)
	defer conn.Close()
	req := &greetpb.GreetManyRequest{
		Greeting: &greetpb.GreetingMany{
			FirstName: "Dlo",
			LastName:  "Bagari",
		},
	}
	stream, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	// read until EOF
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message.GetResult())
	}

}
