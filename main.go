package main

import (
    "flag"
    "fmt"
)

func main() {
    // コマンドライン引数を解析
    params := SearchParams{}
    flag.IntVar(&params.Offset, "offset", 0, "Offset for search results")
    flag.IntVar(&params.Limit, "limit", 10, "Number of recipes per page")
    flag.StringVar(&params.Query, "query", "", "Search query")
    flag.StringVar(&params.Models, "models", "KN-HW24G", "Hotcook model (default: KN-HW24G)")
    flag.StringVar(&params.CookTime, "cooktime", "", "Cooking time")
    flag.BoolVar(&params.Reservation, "reservation", false, "Reservation capable")

    var outputFile string
    flag.StringVar(&outputFile, "output", "", "Output file for JSON data")

    flag.Parse()

    // 検索実行
    htmlContent, err := PerformSearch(&params)
    if err != nil {
        fmt.Println("Error performing search:", err)
        return
    }

    if htmlContent == "" {
        fmt.Println("No HTML content retrieved.")
        return
    }

    // レシピ抽出
    recipes, err := ExtractRecipes(htmlContent)
    if err != nil {
        fmt.Println("Error extracting recipes:", err)
        return
    }

    if len(recipes) == 0 {
        fmt.Println("No recipes found.")
        return
    }

    fmt.Printf("Number of recipes found: %d\n", len(recipes))

    // JSON変換
    jsonOutput, err := RecipesToJSON(recipes)
    if err != nil {
        fmt.Println("Error converting recipes to JSON:", err)
        return
    }

    // JSONファイル出力
    if outputFile != "" {
    err = UpdateJSONFile(recipes, outputFile)
    if err != nil {
        fmt.Println("Error updating JSON file:", err)
        return
    }

    fmt.Printf("JSON data updated in %s\n", outputFile)

    } else {
        fmt.Println("JSON Output:", jsonOutput)
    }

}

