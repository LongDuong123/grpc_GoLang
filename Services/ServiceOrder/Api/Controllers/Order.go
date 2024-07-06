package controllers

import (
	"encoding/json"
	domain "grpc_project/Services/ServiceOrder/Domain"
	"net/http"
)

type OrderController struct {
	OrderInteractor domain.OrderInteractor
}

func NewOrderController(ord domain.OrderInteractor) *OrderController {
	return &OrderController{OrderInteractor: ord}
}

func (ordc *OrderController) UserOrder(w http.ResponseWriter, r *http.Request) {
	var Order domain.Order
	err := json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bill, err := ordc.OrderInteractor.CreateOrder(&Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bill)
}
