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
	mockDB *sql.DB
	err    error
)

//這邊可寫可不寫

// NOTE: golang test支援執行TestMain函數，但也可不寫喔!
func setup() {
	// NOTE: 創建 sqlmock 實例，原本只有sqlmock.New()，但有看到好簡篇教學都說要加上這啥QueryMatcher
	SetMock()

	//NOTE: 不加這會報錯! 意義不明
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

func SetMock() {
	//QueryMatcherEqual -> 這可寫可不寫，用途就是告訴sqlMock你想要怎樣去匹配你的actulSQL和expectedSQL
	// mockDB, Mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//預設是 QueryMatcherRegexp -> 用正規表達去匹配，比Equal匹配更靈活，這也是為啥這是預設值
	mockDB, Mock, err = sqlmock.New()
	if err != nil {
		panic("couldn't create mock: " + err.Error())
	}

}
func teardown() {
	defer mockDB.Close()
}

// NOTE: 测试用的参数有且只有一个 -> t *testing.T。
// NOTE: 基準测试(benchmark)的参数是 *testing.B
// NOTE: TestMain 的参数是 *testing.M 类型。
func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}
