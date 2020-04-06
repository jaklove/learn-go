package main

import (
	"context"
	"google.golang.org/grpc"
	"learn/learn-17"
	"log"
	"os"
	"time"
)

const (
	address = "localhost:50001"
	defaultName = "world"
)

func main()  {
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	if e != nil{
		log.Fatalf("did not connect : %v",e)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)


	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

