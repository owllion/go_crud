package main

import (
	"log"
	"practice/gRPC/server"
	"practice/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

		// launch GIN RESTful API Server
		go func() {
			//one of the servers has to run in a goroutine because it is blocking
			server := router.Setup_Router()
			log.Println("Gin server started on :8080")
			server.Run(":8080")
		}()

		//launch gRPC server
		server.StartGRPCServer()
		
}

