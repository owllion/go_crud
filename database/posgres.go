package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPosgres() {

    dsn := "host=127.0.0.1 user=postgres password=123456789 dbname=education port=5432 sslmode=disable TimeZone=Asia/Shanghai"


	var err error

    PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect to pos")
    }

    fmt.Println("Connected to the pos!")
    
}