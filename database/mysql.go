package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/mysql_student?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to MySQL database failed:", err)
	}

	fmt.Println("Connect to MySQL!")

	// MysqlDB.AutoMigrate(&student.Student{}, &student.Course{}, &student.StudentCourse{})
}