package action

import (
	"context"
	"practice/util"

	"github.com/chromedp/chromedp"
)


func GetPageTitle(ctx context.Context) (title string) {
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.allrecipes.com/"), //導向該網頁
		chromedp.WaitVisible(`body`, chromedp.ByQuery), //等待可見
		chromedp.Title(&title), //獲取標題
	); err != nil {
		util.Log("執行chromedp失敗", "", err.Error())
		title = ""
	}
	util.Log("執行chromedp成功", "", "")
	return 

}
