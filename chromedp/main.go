package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"practice/util"

	"github.com/chromedp/chromedp"
)

func main() {
	// 创建上下文和取消函数
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		return
	}

	// 创建日志文件目录
	logDir := filepath.Join(currentDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Println("Failed to create log directory:", err)
		return
	}

	fmt.Println("Log directory:", logDir)


	// 执行 Chromedp 任务
	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.rakuten.com.tw/"),     // 打开网页
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 等待页面可见
		chromedp.Title(&title),                       // 获取页面标题
	); err != nil {
		util.Log("執行chromedp失敗", "", err.Error())
		return
	}
	util.Log("執行chromedp成功", "" , "")
	fmt.Printf("Page Title: %s\n", title)
}
