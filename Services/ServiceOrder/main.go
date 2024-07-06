package main

import (
	routes "grpc_project/Services/ServiceOrder/Api/Routes"
	"grpc_project/Services/ServiceOrder/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	UseCase, err := config.InitServiceOrder()
	if err != nil {
		log.Fatal("Fail connect Server", err)
	}
	r := mux.NewRouter()
	routes.RegisterRouterOrder(r, UseCase.OrderUseCase)
	http.Handle("/", r)
	http.ListenAndServe(":8081", (r))
}
