package main

import (
	controllers "grpc_project/Services/ServiceBooks/Api/Controllers"
	usecase "grpc_project/Services/ServiceBooks/Application/UseCase"
	"grpc_project/Services/ServiceBooks/Application/proto"
	infrastructure "grpc_project/Services/ServiceBooks/Infrastructure"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	database, err := infrastructure.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	Repository := infrastructure.NewRepositoryBook(database)
	UseCase := usecase.NewBookInteractor(Repository)
	grpcServer := grpc.NewServer()
	proto.RegisterBookServiceServer(grpcServer, &controllers.BookServer{BookUseCase: UseCase})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
