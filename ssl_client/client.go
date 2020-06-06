package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/greet/greetpb"
	"log"
)

func main() {
	// create a insecure connection
	certFile := "ssl_files/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatal(sslErr)
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
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
	// set deadline
	// Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Result)
}
