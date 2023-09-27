package main

import (
	"log"
	"net"

	"practice/gRPC/service/search"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	search.RegisterSearchServiceServer(grpcServer, new(search.SearchServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}