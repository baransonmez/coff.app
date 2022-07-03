package main

import (
	"fmt"
	userGrpc "github.com/baransonmez/coff.app/app/web/grpc/user"
	"github.com/baransonmez/coff.app/app/web/grpc/user/pb"
	coffeeHandlers "github.com/baransonmez/coff.app/app/web/handlers/coffee"
	recipeHandlers "github.com/baransonmez/coff.app/app/web/handlers/recipe"
	userHandlers "github.com/baransonmez/coff.app/app/web/handlers/user"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/output_adapters/persistence"
	"github.com/baransonmez/coff.app/business/core/recipe"
	userClientGrpc "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/grpc"
	recipeData "github.com/baransonmez/coff.app/business/core/recipe/output_adapters/persistence"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

	coffStore := coffeeData.NewInMem()
	coffeeApi := coffeeHandlers.Handlers{
		CoffeeService: coffee.NewService(coffStore),
	}

	userStore := userData.NewInMem()
	userApi := userHandlers.Handlers{UserService: user.NewService(userStore)}

	connAddress := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", connAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, &userGrpc.Server{UserService: user.NewService(userStore)})
	go grpcServer.Serve(lis)

	conn, err := grpc.Dial(connAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	recipeStore := recipeData.NewInMem()
	userClient := userClientGrpc.NewClient(conn)
	recipeApi := recipeHandlers.Handlers{RecipeService: recipe.NewService(recipeStore, userClient)}

	mux := createMux(coffeeApi, recipeApi, userApi)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createMux(coffeeApi coffeeHandlers.Handlers, recipeApi recipeHandlers.Handlers, userApi userHandlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/bean", web.Handle(coffeeApi.Create))
	mux.HandleFunc("/bean/", web.Handle(coffeeApi.GetCoffee))
	mux.HandleFunc("/recipe", web.Handle(recipeApi.Create))
	mux.HandleFunc("/recipe/", web.Handle(recipeApi.Get))
	mux.HandleFunc("/user", web.Handle(userApi.Create))
	mux.HandleFunc("/user/", web.Handle(userApi.Get))
	return mux
}
