package main

import (
	"context"
	"log"
	action "practice/chromedp/actions"

	"github.com/chromedp/chromedp"
)

func main() {
	//創建上下文 & 取消函數
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf), //??
	)
	defer cancel() //確切是何時會關閉?
	
	// //NOTE: 標題
	// title := action.GetPageTitle(ctx)
	// fmt.Println("頁面標題------", title)

	// //NOTE: 內容
	// action.GetHtml(ctx)

	//螢幕截圖
	// action.GetScreenshots(ctx)
	//取得一個食譜類別頁面的html
	action.GetRecipeLinks(ctx)
}

