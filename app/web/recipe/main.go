package main

import (
	"fmt"
	api "github.com/baransonmez/coff.app/app/web/recipe/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/recipe"
	userClientGrpc "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/grpc"
	recipeData "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

	connAddress := "0.0.0.0:50051"
	conn, err := grpc.Dial(connAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	recipeStore := recipeData.NewInMem()
	userClient := userClientGrpc.NewClient(conn)
	recipeApi := api.Handlers{RecipeService: recipe.NewService(recipeStore, userClient)}

	mux := createMux(recipeApi)
	log.Fatal(http.ListenAndServe(":8089", mux))
}

func createMux(recipeApi api.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/recipe", web.Handle(recipeApi.Create))
	mux.HandleFunc("/recipe/", web.Handle(recipeApi.Get))
	return mux
}
