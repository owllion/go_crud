package test

import (
	"database/sql"
	"os"
	db "practice/database"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mock   sqlmock.Sqlmock
	mockDB   *sql.DB
	err error
)
//這邊可寫可不寫

func setup() {
	// 創建 sqlmock 實例
	mockDB, Mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("couldn't create mock: " + err.Error())
	}
	Mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.25"))

	// 使用 sqlmock 替換真實的資料庫
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	if err != nil {
		panic("couldn't open mock db: " + err.Error())
	}

	db.SetDB(gormDB)
}

func teardown() {
	defer mockDB.Close()
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}
