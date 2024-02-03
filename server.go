package main

import (
	"practice/router"
	"practice/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// launch GIN RESTful API Server
	// go func() {
	// 	//one of the servers has to run in a goroutine because it is blocking
	// 	server := router.Setup_Router()
	// 	log.Println("Gin server started on :8080")
	// 	server.Run(":8080")
	// }()
	server := router.Setup_Router()
	// log.Println("Gin server started on :8080")
	util.Log("Gin server started on :8080", "", "")
	server.Run(":8080")

	// etcdClient.StartEtcd()
	// launch gRPC server
	// grpcServer.StartGRPCServer()

}
