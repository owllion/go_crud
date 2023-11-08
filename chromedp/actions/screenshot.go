package action

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func GetScreenshots(ctx context.Context) {
    var buf []byte
    if err := chromedp.Run(ctx, ElementScreenshot(mainUrl, "body", &buf)); err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile("body.png", buf, 0o644); err != nil {
        log.Fatal(err)
    }

    if err := chromedp.Run(ctx, FullScreenshot(mainUrl, 90, &buf)); err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile("full.png", buf, 0o644); err != nil {
        log.Fatal(err)
    }

    fmt.Println("screenshots created")
}

func ElementScreenshot(url, sel string, res *[]byte) chromedp.Tasks {

    return chromedp.Tasks{

        chromedp.Navigate(url),
        chromedp.Screenshot(sel, res, chromedp.NodeVisible),
    }
}

func FullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {

    return chromedp.Tasks{

        chromedp.Navigate(url),
        chromedp.FullScreenshot(res, quality),
    }
}