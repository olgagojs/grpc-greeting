package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"../greetpb"
)

func main() {
	fmt.Println("Hello I'm client!");
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure());
	if err != nil {
		log.Fatalf("Could not connect: %v", err);
	}

	defer cc.Close();

	c := greetpb.NewGreetServiceClient(cc);

	doUnary(c);
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an Unary RPC...");
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Ejik",
			LastName: "VSeti",
		},
	}
	res, err := c.Greet(context.Background(), req);
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err);
	}
	log.Printf("Responce from Greet RPC: %f", res.Result);
}