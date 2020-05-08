package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/greet/greetpb"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := greetpb.NewBidiConvClient(conn)
	stream, err := c.Bidi(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	waitc := make(chan struct{})

	go func() {
		values := []string{"client1", "client2", ""}
		for _, i := range values {
			if err := stream.Send(&greetpb.BidiRequest{
				Person: &greetpb.Person{Name: i}}); err != nil {
				log.Fatal(err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			messg, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				break
			}
			// dealing with err. this error can be grpc error or another error
			// first we try to convert the err to grpc error, if the conversion is succeed  we habdle the error, otherwise the error is not grpc error
			if err != nil {
				//convert error to grpc.
				respError, ok := status.FromError(err)
				//if ok is true, then this error is grpc error
				if ok {
					fmt.Println(respError.Message())
					fmt.Println(respError.Code())
					if respError.Code() == codes.InvalidArgument {
						fmt.Println("invalid argument was sent")
					}
				} else {
					//the error is golang error
					log.Println(err)
				}
				close(waitc)
				break
			}
			fmt.Println(messg.GetResult())
		}
	}()
	<-waitc
}
