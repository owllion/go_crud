package main

//NOTE: 要執行測試，只要在有_test.go檔案所在的目錄即可
//NOTE: 但如果被測試函數和測試檔案所在目錄不同，那 -cover 就無法看到覆蓋率喔~
//NOTE: 至於官方則是推薦把被測 & 測試文件都放到同個目錄~
func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}
