package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"log"
)

func main() {
	// create a insecure connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connected with server: %s", err)
	}
	c := greetpb.NewGreetServiceClient(conn)
	defer conn.Close()

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dlo",
			LastName:  "Bagari",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Result)

}
