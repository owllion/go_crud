package db

import (
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	MysqlDB *gorm.DB
	PostgresDB *gorm.DB
)