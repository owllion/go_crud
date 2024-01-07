package main

import (
	"fmt"
	"practice/util"
	"time"

	excelize "github.com/xuri/excelize/v2"
)

//TODO: 左上日期 / 右倒數1 & 2 & 3 合計 / 第一列合併儲存格，動態值

func main() {
	header := "業務"
	data := util.GenFakeData() //拿到的資料

	//設定標題(要合併，長度就是所有欄位的長度，也就是A1~X1)
	//這邊x就會是欄位數量對應的字母


	f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()


	// index, err := f.NewSheet("Sheet1")
	
	//檔案標題
	f.SetCellValue("Sheet1", "A1", fmt.Sprintf("訂單彙總表 - %v",header))
	//合併標題欄位
	f.MergeCell("Sheet1", "A1", "C1")

	//設定欄位標題
	f.SetCellValue("Sheet1", "A2", "Name")
	f.SetCellValue("Sheet1", "B2", "Qty")
	f.SetCellValue("Sheet1", "C2", "Product")

	for r, item := range data {
		f.SetCellValue("Sheet1", "A"+fmt.Sprint(r+3),item.Name)
		f.SetCellValue("Sheet1", "B"+fmt.Sprint(r+3), item.Qty)
    	f.SetCellValue("Sheet1", "C"+fmt.Sprint(r+3), item.Price)

	}

	sheetName := time.Now().Format("2024-01-07")
	path := fmt.Sprintf("../assets/excel/%s.xlsx", sheetName)
	err := f.SaveAs(path)
	if err != nil {
		fmt.Println(err)
	}

	


}
