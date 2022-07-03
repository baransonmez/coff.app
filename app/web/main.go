package main

import (
	"fmt"
	coffeeHandlers "github.com/baransonmez/coff.app/app/web/handlers/coffee"
	recipeHandlers "github.com/baransonmez/coff.app/app/web/handlers/recipe"
	userHandlers "github.com/baransonmez/coff.app/app/web/handlers/user"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/output_adapters/persistence"
	"github.com/baransonmez/coff.app/business/core/recipe"
	recipeData "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/persistence"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"log"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

	coffStore := coffeeData.NewInMem()
	coffeeApi := coffeeHandlers.Handlers{
		CoffeeService: coffee.NewService(coffStore),
	}

	recipeStore := recipeData.NewInMem()
	recipeApi := recipeHandlers.Handlers{RecipeService: recipe.NewService(recipeStore)}

	userStore := userData.NewInMem()
	userApi := userHandlers.Handlers{UserService: user.NewService(userStore)}

	mux := http.NewServeMux()
	mux.HandleFunc("/bean", web.Handle(coffeeApi.Create))
	mux.HandleFunc("/bean/", web.Handle(coffeeApi.GetCoffee))
	mux.HandleFunc("/recipe", web.Handle(recipeApi.Create))
	mux.HandleFunc("/recipe/", web.Handle(recipeApi.Get))
	mux.HandleFunc("/user", web.Handle(userApi.Create))
	mux.HandleFunc("/user/", web.Handle(userApi.Get))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
