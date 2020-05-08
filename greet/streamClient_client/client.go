package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := greetpb.NewGreetClientStreamClient(conn)
	data := []*greetpb.GreetClientRequest{{
		Greeting: &greetpb.GreetClient{
			Name: "Dlo",
		},
	},
		{Greeting: &greetpb.GreetClient{
			Name: "Bagari",
		},
		},
	}
	stream, err := c.Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range data {
		err := stream.Send(i)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetResult())
}
