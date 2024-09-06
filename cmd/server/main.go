package main

import (
	"log"
	"myGRPC/interanal/adder"
	"myGRPC/proto/createdProto"
	"net"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()

	srv := &adder.GRPCServer{}

	createdProto.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
