package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPosgres() {

    dsn := "host=192.168.0.20 user=postgres password=50984878 dbname=01_GZ_SAT_ES port=5432 sslmode=disable TimeZone=Asia/Shanghai"


	var err error

    PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect to pos")
    }

    fmt.Println("Connected to the pos!")
    
}