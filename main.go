package main

import (
	"fmt"
	"google.golang.org/grpc"
	"micro-maps/communicate"
	"micro-maps/controller"
	"net"
)

func main() {

	// port := os.Getenv("PORT")
	port := 7000
	host := fmt.Sprintf("0.0.0.0:%v", port)

	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	communicate.RegisterGelocationCommunicateServer(grpcServer, &controller.MapServer{})

	fmt.Printf("[x] - Server logistic listen http://localhost:%v\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err.Error())
	}
}
