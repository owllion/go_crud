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
	MysqlDB = mockSql
}