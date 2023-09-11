package main

import (
	sql "practice/database"
	"practice/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sql.InitDatabase()
	server := router.Setup_Router()
	server.Run(":5487")
}

