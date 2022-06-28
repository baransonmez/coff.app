package main

import (
	"fmt"
	coffeeHandlers "github.com/baransonmez/coff.app/app/web/handlers/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/data"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/bean", web.Handle(coffeeApi.Create))
	mux.HandleFunc("/bean/", web.Handle(coffeeApi.GetCoffee))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
