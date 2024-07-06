package routes

import (
	controllers "grpc_project/Services/ServiceBooks/Api/Controllers"
	domain "grpc_project/Services/ServiceBooks/Domain"

	"github.com/gorilla/mux"
)

func RegisterRouterBooks(r *mux.Router, bookUseCase domain.BookInteractor) {
	BookControllers := controllers.NewBookController(bookUseCase)
	r.HandleFunc("/book/{id}", BookControllers.GetBook).Methods("GET")
}
