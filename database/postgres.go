package db

import (
	"fmt"
	student "practice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitPostgres() {

    dsn := "host=127.0.0.1 user=postgres password=123456789 dbname=education port=5432 sslmode=disable TimeZone=Asia/Shanghai"


	var err error

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })

    if err != nil {
        panic("failed to connect to pos")
    }

    fmt.Println("Connected to the pos!")

    DB.AutoMigrate(&student.Student{}, &student.Course{}, &student.StudentCourse{})
    
}