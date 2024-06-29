package main

import (
	rabbitmqUseCase "grpc_project/Services/ServiceOrder/Application/RabbitMQ"
	usecase "grpc_project/Services/ServiceOrder/Application/UseCase"
	domain "grpc_project/Services/ServiceOrder/Domain"
	rabbitmq "grpc_project/Services/ServiceOrder/Infrastructure/RabbitMQ"
	grpcbook "grpc_project/Services/ServiceOrder/Infrastructure/grpc_book"
	"log"
)

func main() {
	rabbitMQ, err := rabbitmq.ConnectServiceEmail()
	if err != nil {
		log.Fatal("Fail connect rabbitMQ")
	}
	grpc_book, err := grpcbook.ConnectServiceBooks()
	if err != nil {
		log.Fatal("Fail connect grpc")
	}
	Sendemail := rabbitmqUseCase.NewSendEmail(rabbitMQ)
	Repository := grpcbook.NewRepositoryGprc(grpc_book)
	order := usecase.NewOrderInteractor(Repository, *Sendemail)
	orderByUser := &domain.Order{UserID: 1, BookID: []int64{1, 2, 3}, AmountRequired: []int64{2, 2, 2}, Total: 0}
	orderProcessing, err := order.CreateOrder(orderByUser)
	if err != nil {
		log.Fatal("Fail orderProcessing :", err)
	}
	log.Println(orderProcessing)
}
