package main

import (
	userGrpc "github.com/baransonmez/coff.app/app/web/user/input_adapters/grpc"
	"github.com/baransonmez/coff.app/app/web/user/input_adapters/grpc/pb"
	api "github.com/baransonmez/coff.app/app/web/user/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	userStore := userData.NewInMem()
	userApi := api.Handlers{UserService: user.NewService(userStore)}

	connAddress := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", connAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, &userGrpc.Server{UserService: user.NewService(userStore)})
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("grpc server does not start!")
		}
	}()

	handler := routes(userApi)
	servPort := ":8080"
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
	router.HandlerFunc(http.MethodPost, "/v1/user", web.Handle(recipeApi.Create))
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", web.Handle(recipeApi.Get))
	return router
}
