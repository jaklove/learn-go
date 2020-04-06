package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn/learn-17"
	"log"
	"net"
)

const(
	PORT = ":50001"
)

type server struct {}

func (s *server)SayHello (ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply,error){
	fmt.Println("request:",in.Name)
	return &helloworld.HelloReply{Message:"Hello "+ in.Name},nil
}

func main()  {
	listener, e := net.Listen("tcp", PORT)
	if e != nil{
		log.Fatalf("failed to listen: %v",e)
	}
	newServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(newServer,&server{})
	log.Println("rpc服务已经开启")
	newServer.Serve(listener)
}


