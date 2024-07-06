package routes

import (
	controllers "grpc_project/Services/ServiceOrder/Api/Controllers"
	domain "grpc_project/Services/ServiceOrder/Domain"

	"github.com/gorilla/mux"
)

func RegisterRouterOrder(r *mux.Router, OrderInterator domain.OrderInteractor) {
	orderController := controllers.NewOrderController(OrderInterator)
	r.HandleFunc("/order", orderController.UserOrder)
}
