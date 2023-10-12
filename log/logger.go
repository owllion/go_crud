package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)


func SetLogger() (myLogger logger.Interface) {
	t := time.Now().Local().Format("2006-01-02")

	fileName := fmt.Sprint("./log/", t, ".log")

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("file open error : %v", err)
	}

	myLogger = logger.New(log.New(file, "", log.LstdFlags), // io writer

		// LogLevel 定义了应该记录哪些日志消息。可以是 Silent、Error、Warn 或 Info。
		// Silent: 不记录任何日志
		// Error: 只记录错误事件
		// Warn: 记录警告和错误事件
		// Info: 记录所有日志事件（包括查询、警告和错误）
		logger.Config{
			 // SlowThreshold 设定了多久的 SQL 查询应被认为是 "slow" 或较慢的查询。对于超过这个时间阈值的查询，日志中会特别标记。
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, 
			Colorful:      true,  
		},
	)

	return
}