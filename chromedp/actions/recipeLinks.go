package action

import (
	"context"
	"fmt"
	"practice/util"

	"github.com/chromedp/chromedp"
)

func GetRecipeLinks(ctx context.Context) (data string) {

	selector := "#header-nav_1-0 > div.header-nav__list-wrapper > ul > li:nth-child(1) > ul > li:nth-child(1) > a"


	if err := chromedp.Run(
		ctx,
        chromedp.Navigate(mainUrl),
		
		chromedp.Click(selector, chromedp.ByQuery),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
		); err != nil {
			util.Log("chromedp.Nodes 執行失敗", "", err.Error())
			return 
		}
	fmt.Println("data", data)
	return 


}