package main

import (
	"encoding/json"
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/data/persistence"
	"github.com/baransonmez/coff.app/business/core/recipe"
	recipeData "github.com/baransonmez/coff.app/business/core/recipe/data/persistence"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/data/persistence"
	"time"
)

func main() {
	fmt.Println("mod file generated")
	coffStore := coffeeData.NewInMem()
	service := coffee.NewService(coffStore)
	beanId, _ := service.CreateCoffeeBean(nil, coffee.NewCoffeeBean{
		Name:      "Yirgaciffe",
		Roaster:   "Montag",
		Origin:    "Etiopia",
		Price:     23,
		RoastDate: time.Now().AddDate(2, 3, 4),
	})
	bean, _ := service.GetCoffeeBean(nil, beanId)
	fmt.Println(prettyPrint(bean))

	userStore := userData.NewInMem()
	userService := user.NewService(userStore)
	userId, _ := userService.CreateNewUser(nil, user.NewUser{
		Name: "Baran",
	})
	newUser, _ := userService.GetUser(nil, userId)
	fmt.Println(prettyPrint(newUser))

	recipeStore := recipeData.NewInMem()
	recipeService := recipe.NewService(recipeStore)
	recipeId, _ := recipeService.CreateNewRecipe(nil, recipe.NewRecipe{
		UserID:      userId.String(),
		CoffeeID:    beanId.String(),
		Description: "30 seconds blooming",
		Steps: []recipe.Step{
			{
				Description:       "blooming",
				DurationInSeconds: 24,
			},
			{
				Description:       "brewing",
				DurationInSeconds: 76,
			}},
	})

	newRecipe, _ := recipeService.GetRecipe(nil, recipeId)
	fmt.Println(prettyPrint(newRecipe))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
