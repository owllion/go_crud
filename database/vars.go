package db

import (
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	MysqlDB *gorm.DB
	PostgresDB *gorm.DB
)


//測試用: mock mysql database
//使用方法: 測試檔案中手動呼叫
func SetDB(mockSql *gorm.DB) {
	DB = mockSql
	PostgresDB = mockSql
}
/* NOTE: 
db賦值流程:
	1.有import就會跑init & go test -v 會最先執行TestMain function -> 那寫在 main_test.go裡面 ->所以main_test會先被讀取 -> main_test裡有引入db -> 因此會先正常初始化一次(應該有，不確定) -> 裡面又有執行 db.SetDb(建立的mock db) -> 目前db就是mock db
	2.換成student_test執行 -> 有引入GetStudent controller -> 這裡面有引入db -> 但 init已經執行過，不會再執行 -> 因此當下的db就會是 mock db.
*/