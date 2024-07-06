package config

import (
	usecase "grpc_project/Services/ServiceBooks/Application/UseCase"

	domain "grpc_project/Services/ServiceBooks/Domain"
	infrastructure "grpc_project/Services/ServiceBooks/Infrastructure"
	"net"
)

type Init struct {
	Li          net.Listener
	UseCaseBook domain.BookInteractor
}

func InitServer() (*Init, error) {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}
	database, err := infrastructure.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	Repository := infrastructure.NewRepositoryBook(database)
	UseCase := usecase.NewBookInteractor(Repository)
	return &Init{Li: lis, UseCaseBook: UseCase}, nil
}
