package main

import (
	controllers "grpc_project/Services/ServiceBooks/Api/Controllers"
	routes "grpc_project/Services/ServiceBooks/Api/Routes"
	"grpc_project/Services/ServiceBooks/Application/proto"
	"grpc_project/Services/ServiceBooks/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	Init, err := config.InitServer()
	if err != nil {
		log.Fatal("Fail Connect Server", err)
		return
	}
	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterBookServiceServer(grpcServer, &controllers.BookServer{BookUseCase: Init.UseCaseBook})
		if err := grpcServer.Serve(Init.Li); err != nil {
			log.Fatal("Fail Server", err)
		}
	}()
	r := mux.NewRouter()
	routes.RegisterRouterBooks(r, Init.UseCaseBook)
	http.Handle("/", r)
	http.ListenAndServe(":8084", (r))
}
