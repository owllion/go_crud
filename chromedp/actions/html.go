package action

import (
	"context"
	"fmt"
	"practice/util"

	"github.com/chromedp/chromedp"
)


func GetHtml(ctx context.Context) {
	
	// url := "https://www.allrecipes.com/"
	// if err := chromedp.Run(ctx,
	// 	chromedp.Navigate(url),
	// ); err != nil {
	// 	util.Log("getRecipes - 執行chromedp失敗", "", err.Error())
	// }


	// util.Log("getRecipes - 導向網頁成功", "", "")
	// var recipes []*cdp.Node
	// if err := chromedp.Run(ctx,
	// 	chromedp.Nodes("div", &recipes, chromedp.NodeVisible),
	// ); err != nil {
	// 	util.Log("chromedp.Nodes 執行失敗", "", err.Error())
	// 	return
	// }
	// util.Log("getRecipes - chromedp.Nodes 執行成功", "", "")

	// for idx, recipe := range recipes {
	// 	fmt.Printf("---------------------------食譜%d: %v",idx,recipe)
	// }

    var data string

    if err := chromedp.Run(ctx,
        chromedp.Navigate(mainUrl),
        chromedp.OuterHTML("html", &data, chromedp.ByQuery),
    ); err != nil {
		util.Log("chromedp.Nodes 執行失敗", "", err.Error())
    }

    fmt.Println(data)
}
