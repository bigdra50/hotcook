package main

import (
    "encoding/json"
    "golang.org/x/net/html"
    "strings"
)

func ExtractRecipes(htmlContent string) ([]Recipe, error) {
    doc, err := html.Parse(strings.NewReader(htmlContent))
    if err != nil {
        return nil, err
    }

    var recipes []Recipe
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "img" {
            var recipe Recipe
            for _, a := range n.Attr {
                if a.Key == "src" && strings.Contains(a.Val, "/kitchen/recipe/photo/") {
                    parts := strings.Split(a.Val, "/")
                    if len(parts) > 4 {
                        recipe.ID = strings.TrimSuffix(parts[4], ".jpg")
                    }
                }
                if a.Key == "alt" {
                    recipe.Name = a.Val
                }
            }
            if recipe.ID != "" && recipe.Name != "" {
                recipes = append(recipes, recipe)
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)

    return recipes, nil
}

func RecipesToJSON(recipes []Recipe) (string, error) {
    jsonData, err := json.Marshal(recipes)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}

