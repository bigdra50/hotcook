package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "fmt"
)


// UpdateJSONFile updates or appends new recipes in the specified JSON file.
func UpdateJSONFile(recipes []Recipe, filename string) error {
    var existingRecipes []Recipe

    // Check if file exists
    if _, err := os.Stat(filename); err == nil {
        fileContent, err := ioutil.ReadFile(filename)
        if err != nil {
            return err
        }

        err = json.Unmarshal(fileContent, &existingRecipes)
        if err != nil {
            return err
        }
    }

    // Update or append recipes
    for _, newRecipe := range recipes {
        newRecipe.ImageURL = fmt.Sprintf("https://cocoroplus.jp.sharp/kitchen/recipe/photo/%s.jpg", newRecipe.ID)
        updated := false
        for i, existingRecipe := range existingRecipes {
            if existingRecipe.ID == newRecipe.ID {
                existingRecipe.ImageURL = newRecipe.ImageURL
                existingRecipes[i] = newRecipe
                updated = true
                break
            }
        }
        if !updated {
            existingRecipes = append(existingRecipes, newRecipe)
        }
    }

    updatedContent, err := json.MarshalIndent(existingRecipes, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filename, updatedContent, 0644)
}
