package action

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func GetRecipe(ctx context.Context) {
	
	selector := "#header-nav_1-0 > div.header-nav__list-wrapper > ul > li:nth-child(1) > ul > li:nth-child(1) > a"

	sel := "document.querySelector(`body`)"

    var htmlContent string
	err := chromedp.Run(ctx,
        chromedp.Navigate(mainUrl),
        chromedp.WaitVisible(selector),
        chromedp.Click(selector),
    )
    if err != nil {
        log.Fatal(err)
    }
	
	title := GetPageTitle(ctx)
	fmt.Println("葉面標題", title)
	 // 等待新页面加载完成
	 if err := chromedp.WaitReady("html"); err != nil {
        log.Fatal(err)
    }

    // 获取点击后页面的 HTML 内容
    err = chromedp.Run(ctx,
        chromedp.OuterHTML(sel, &htmlContent, chromedp.ByQuery),
    )
    if err != nil {
        log.Fatal(err)
    }

    // 打印获取的 HTML 内容
    fmt.Println("Page HTML:", htmlContent)



    dropdownOptionSelector := "#header-nav_1-0 > div.header-nav__list-wrapper > ul > li:nth-child(1) > ul > li:nth-child(1) > a" 

    // 点击下拉菜单中的选项
    if err := chromedp.Run(ctx, chromedp.Click(dropdownOptionSelector)); err != nil {
        log.Fatal(err)
    }

	// 获取两个不同的div区块的选择器
	// firstDivSelector := "#three-post__inner_1-0" // 第一个div区块的选择器
	// secondDivSelector := "#mntl-taxonomysc-article-list-group_1-0" // 第二个div区块的选择器

	// 获取第一个div区块中的食谱
	// var recipeNodes1 []*chromedp.Node
	// if err := chromedp.Nodes(ctx, firstDivSelector+" a[id^='mntl-card-list-items_2-0-']", &recipeNodes1); err != nil {
	// 	log.Fatal(err)
	// }

	// // 遍历第一个div区块，处理每个食谱的内容
	// for i, recipeNode := range recipeNodes1 {
	// 	// 获取每个食谱的文本内容
	// 	recipeText := ""
	// 	if err := chromedp.Run(ctx, chromedp.Text(recipeNode, &recipeText)); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// 打印每个食谱的文本内容
	// 	fmt.Printf("Recipe %d:\n%s\n", i+1, recipeText)
	// }

	// // 获取第二个div区块中的食谱
	// var recipeNodes2 []*chromedp.Node
	// if err := chromedp.Nodes(ctx, secondDivSelector+" div[id^='tax-sc__recirc-list-container_1-0'] a[id^='mntl-card-list-items_2-0-']", &recipeNodes2); err != nil {
	// 	log.Fatal(err)
	// }

	// // 遍历第二个div区块，处理每个食谱的内容
	// for i, recipeNode := range recipeNodes2 {
	// 	// 获取每个食谱的文本内容
	// 	recipeText := ""
	// 	if err := chromedp.Run(ctx, chromedp.Text(recipeNode, &recipeText)); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// 打印每个食谱的文本内容
	// 	fmt.Printf("Recipe %d:\n%s\n", i+1, recipeText)
	// }
}
	
