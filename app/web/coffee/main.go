package main

import (
	"fmt"
	api "github.com/baransonmez/coff.app/app/web/coffee/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"log"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

	coffStore := coffeeData.NewInMem()
	coffeeApi := api.Handlers{
		CoffeeService: coffee.NewService(coffStore),
	}

	mux := createMux(coffeeApi)
	log.Fatal(http.ListenAndServe(":8085", mux))
}

func createMux(coffeeApi api.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/bean", web.Handle(coffeeApi.Create))
	mux.HandleFunc("/bean/", web.Handle(coffeeApi.GetCoffee))
	return mux
}
