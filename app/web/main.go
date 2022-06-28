package main

import (
	"fmt"
	coffeeHandlers "github.com/baransonmez/coff.app/app/web/handlers/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/data"
	"log"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

	coffStore := coffeeData.NewInMem()
	pgh := coffeeHandlers.Handlers{
		CoffeeService: coffee.NewService(coffStore),
	}

	create := func(w http.ResponseWriter, r *http.Request) {
		err := pgh.Create(w, r)
		if err != nil {
			return
		}
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		err := pgh.GetCoffee(w, r)
		if err != nil {
			return
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/getBean/", get)
	mux.HandleFunc("/bean", create)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
