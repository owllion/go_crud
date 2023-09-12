package sql

import (
	"fmt"
	"practice/models/student"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to SQLite database failed:", err)
		return
	}
	fmt.Println(DB)
	DB.AutoMigrate(&student.Student{})
}


