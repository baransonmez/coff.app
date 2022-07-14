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

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	//recipeStore := recipeData.NewInMem()
	recipeStore, err := recipeData.NewStore()
	if err != nil {
		log.Fatalln(err)
	}

	userClient := userClientGrpc.NewClient(conn)
	recipeAPI := api.Handlers{RecipeService: recipe.NewService(recipeStore, userClient)}

	handler := routes(recipeAPI)

	servPort := ":8089"
	log.Printf("starting server on %s\n", servPort)

	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func routes(recipeAPI api.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/recipe", web.Handle(recipeAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/recipe/:id", web.Handle(recipeAPI.Get))
	return router
}
