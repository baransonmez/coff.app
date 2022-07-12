package main

import (
	api "github.com/baransonmez/coff.app/app/web/coffee/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	coffStore := coffeeData.NewInMem()
	coffeeAPI := api.Handlers{
		CoffeeService: coffee.NewService(coffStore),
	}

	handler := routes(coffeeAPI)

	servPort := ":8085"
	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}

	log.Printf("starting server on %s\n", servPort)

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func routes(beanAPI api.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/bean", web.Handle(beanAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/bean/:id", web.Handle(beanAPI.GetCoffee))
	return router
}
