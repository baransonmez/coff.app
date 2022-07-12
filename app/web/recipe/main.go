package main

import (
	api "github.com/baransonmez/coff.app/app/web/recipe/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/recipe"
	userClientGrpc "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/grpc"
	recipeData "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func main() {

	connAddress := "0.0.0.0:50051"
	conn, err := grpc.Dial(connAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	recipeStore := recipeData.NewInMem()
	userClient := userClientGrpc.NewClient(conn)
	recipeApi := api.Handlers{RecipeService: recipe.NewService(recipeStore, userClient)}

	handler := routes(recipeApi)

	servPort := ":8089"
	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}

	log.Printf("starting server on %s\n", servPort)
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func routes(recipeApi api.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/recipe", web.Handle(recipeApi.Create))
	router.HandlerFunc(http.MethodGet, "/v1/recipe/:id", web.Handle(recipeApi.Get))
	return router
}
