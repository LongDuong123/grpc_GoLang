package controllers

import (
	"encoding/json"
	"grpc_project/Services/ServiceBooks/Application/proto"
	domain "grpc_project/Services/ServiceBooks/Domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookControllers struct {
	bookInteractor domain.BookInteractor
}

func NewBookController(bki domain.BookInteractor) *BookControllers {
	return &BookControllers{bookInteractor: bki}
}

func (bkc *BookControllers) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdBook, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	IdBookProto := &proto.BookID{Id: IdBook}
	Book, err := bkc.bookInteractor.GetBookByID(IdBookProto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Book)
}
