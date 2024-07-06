package config

import (
	rabbitmqUseCase "grpc_project/Services/ServiceOrder/Application/RabbitMQ"
	usecase "grpc_project/Services/ServiceOrder/Application/UseCase"
	domain "grpc_project/Services/ServiceOrder/Domain"
	postgresql "grpc_project/Services/ServiceOrder/Infrastructure/PostgreSQL"
	rabbitmq "grpc_project/Services/ServiceOrder/Infrastructure/RabbitMQ"
	grpcbook "grpc_project/Services/ServiceOrder/Infrastructure/grpc_book"
)

type ServerOrder struct {
	OrderUseCase domain.OrderInteractor
}

func InitServiceOrder() (*ServerOrder, error) {
	rabbitMQ, err := rabbitmq.ConnectServiceEmail()
	if err != nil {
		return nil, err
	}
	grpc_book, err := grpcbook.ConnectServiceBooks()
	if err != nil {
		return nil, err
	}
	Postgresql, err := postgresql.ConnectPostgreSQL()
	if err != nil {
		return nil, err
	}
	Sendemail := rabbitmqUseCase.NewSendEmail(rabbitMQ)
	GrpcRepository := grpcbook.NewRepositoryGprc(grpc_book)
	PostgreRepository := postgresql.NewRepositoryOrder(Postgresql)
	orderusecase := usecase.NewOrderInteractor(GrpcRepository, *Sendemail, PostgreRepository)
	return &ServerOrder{OrderUseCase: orderusecase}, nil
}
