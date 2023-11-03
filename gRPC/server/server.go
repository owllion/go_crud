package grpcServer

import (
	"fmt"
	"log"
	"net"

	"practice/gRPC/client"
	"practice/gRPC/service/search"

	"google.golang.org/grpc"
)

func StartGRPCServer() {
	grpcServer := grpc.NewServer()

	search.RegisterSearchServiceServer(grpcServer, new(search.SearchServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	client.InitGRPCClient()
	defer client.ClientConn.Close()
	//在gRPC服務停止後才會被調用 

	fmt.Println("gRPC server on port 1234 is running")
	//不能寫在Serve後面 因為到那行就會持續執行，不會再往下走了(大概)
	
	grpcServer.Serve(lis)
	
}
