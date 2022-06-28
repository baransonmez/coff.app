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

	h := func(w http.ResponseWriter, r *http.Request) {
		err := pgh.Create(w, r)
		if err != nil {
			return
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/bean", h)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
