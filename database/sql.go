package sql

import (
	"fmt"
	student "practice/models/student"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to SQLite database failed:", err)
		return
	}
	DB.AutoMigrate(&student.Student{})
}


