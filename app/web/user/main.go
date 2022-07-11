package main

import (
	"fmt"
	userGrpc "github.com/baransonmez/coff.app/app/web/user/input_adapters/grpc"
	"github.com/baransonmez/coff.app/app/web/user/input_adapters/grpc/pb"
	api "github.com/baransonmez/coff.app/app/web/user/input_adapters/handlers"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/output_adapters/persistence"
	"github.com/baransonmez/coff.app/foundation/web"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("web api generated")

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

	mux := createMux(userApi)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createMux(userApi api.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", web.Handle(userApi.Create))
	mux.HandleFunc("/user/", web.Handle(userApi.Get))
	return mux
}
