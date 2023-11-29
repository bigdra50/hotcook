package main

import (
    "testing"
)

func TestExtractRecipes(t *testing.T) {
    htmlContent := `<img src="/kitchen/recipe/photo/R4012.jpg" alt="おでん">`

    expected := []Recipe{{ID: "R4012", Name: "おでん"}}
    recipes, err := ExtractRecipes(htmlContent)

    if err != nil {
        t.Fatalf("Error extracting recipes: %v", err)
    }

    if len(recipes) != len(expected) {
        t.Fatalf("Expected %d recipes, got %d", len(expected), len(recipes))
    }

    for i, recipe := range recipes {
        if recipe.ID != expected[i].ID || recipe.Name != expected[i].Name {
            t.Errorf("Expected %v, got %v", expected[i], recipe)
        }
    }
}

