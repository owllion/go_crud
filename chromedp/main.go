package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {
	// 创建上下文和取消函数
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 执行 Chromedp 任务
	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.rakuten.com.tw/"),     // 打开网页
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 等待页面可见
		chromedp.Title(&title),                       // 获取页面标题
	); err != nil {
		util.Lo
	}

	fmt.Printf("Page Title: %s\n", title)
}
