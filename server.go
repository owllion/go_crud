package main

import (
	"practice/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server := router.Setup_Router()
	server.Run(":5487")
}

