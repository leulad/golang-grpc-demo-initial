package main

import (
	"context"
	"log"
	"time"

	pb "github.com/leulad/golang-grpc-demo-initial/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("%s", res.Message)
}
