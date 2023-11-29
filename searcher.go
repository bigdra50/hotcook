package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    "net/url"
    "time"
    "strings"
)

type SearchParams struct {
    Offset      int
    Limit       int
    Query       string
    Models      string
    CookTime    string
    Reservation bool
}

func buildRequestURL(params *SearchParams) string {
    baseURL := "https://cocoroplus.jp.sharp/kitchen/recipe/searchresults/"
    queryParams := url.Values{}
    queryParams.Set("offset", fmt.Sprintf("%d", params.Offset))
    queryParams.Set("limit", fmt.Sprintf("%d", params.Limit))
    queryParams.Set("search", params.Query)
    queryParams.Set("models", params.Models)
    queryParams.Set("cooktime", params.CookTime)
    queryParams.Set("reservation", fmt.Sprintf("%t", params.Reservation))
    return baseURL + "?" + queryParams.Encode()
}

func PerformSearch(params *SearchParams) (string, error) {
    requestURL := buildRequestURL(params)

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 200*time.Second)
    defer cancel()

    var htmlContent, resultsText string
    var prevScrollHeight, currScrollHeight int64

    err := chromedp.Run(ctx,
        chromedp.Navigate(requestURL),
        chromedp.WaitVisible(`body`, chromedp.ByQuery),
        chromedp.WaitVisible(`.DefaultText.searchResults_num`, chromedp.ByQuery),
        chromedp.Sleep(1*time.Second),
        chromedp.Text(`.DefaultText.searchResults_num`, &resultsText, chromedp.ByQuery),
        chromedp.Evaluate(`document.body.scrollHeight`, &prevScrollHeight),
    )
    if err != nil {
        return "", err
    }

    fmt.Println("Current number of recipes:", resultsText)
    for {
        err := chromedp.Run(ctx,
            chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
            chromedp.Sleep(1*time.Second),
            chromedp.Text(`.DefaultText.searchResults_num`, &resultsText, chromedp.ByQuery),
            chromedp.Evaluate(`document.body.scrollHeight`, &currScrollHeight),
        )
        if err != nil {
            return "", err
        }

        if currScrollHeight == prevScrollHeight {
            break
        }
        prevScrollHeight = currScrollHeight
        fmt.Println("Scrolling...")
    }

    err = chromedp.Run(ctx,
        chromedp.Text(`.DefaultText.searchResults_num`, &resultsText, chromedp.ByQuery),
        chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
    )
    if err != nil {
        return "", err
    }

    // 検索結果が0件の場合のチェック
    if strings.Contains(resultsText, "0件") {
        return "", fmt.Errorf("no recipes found")
    }

    return htmlContent, nil
}
