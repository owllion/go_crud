package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	log.SetPrefix("【Debug】")
	log.SetFlags(log.Ldate | log.Ltime)
	//NOTE: 會長這樣-> 【Debug】2023/10/23 13:05:08 獲取失敗
}

// NOTE: 實際使用: util.Log("獲取失敗", nil, result.Error.Error())
func Log(information string, query string, error string) {
	currentTimeStr := time.Now().Local().Format("2006-01-02")

	fileName := fmt.Sprintf("./log/%s.log", currentTimeStr)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	/*
		NOTE:
			os.O_RDWR 表示以可讀寫方式打開文件
			os.O_CREATE 表示如果文件不存在則創建它
			os.O_APPEND 表示以附加數據的方式打開文件
	*/

	if err != nil {
		log.Fatalf("file open error : %v", err)
		return
	}

	//NOTE: 確保錯誤內容寫入後，文件有正確關閉，避免資源洩漏
	defer file.Close()

	//NOTE: 把log輸出位置設定為上面打開的file
	log.SetOutput(file)

	//NOTE: 這三行並不是在終端印出，而是會寫入log file中
	log.Printf("Info----- %+v\n", information)
	//printf和sprintf基本功能一樣，差在前者只是印出，後者會回傳值
	/*
		NOTE: 佔位符
			%s：用於字符串。
			%d：用於有符號整數。
			%f：用於浮點數。
			%t：用於布爾值。
			%v：用於通用的值，以默認方式格式化。
			%T：用於顯示值的類型。
	*/
	log.Printf("query----- %+v\n", query)
	log.Printf("error----- %+v\n", error)
	log.Println("------------------------------------------")

}

/*NOTE:log會像這樣
【Debug】2023/10/23 13:07:48 獲取失敗
【Debug】2023/10/23 13:07:48 <nil>
【Debug】2023/10/23 13:07:48 ERROR: relation "9WHM_1RUN.runz0wheel" does not exist (SQLSTATE 42P01)
*/
