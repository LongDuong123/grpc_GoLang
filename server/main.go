package main

import (
	"context"
	"log"
	"net"

	pb "grpc_project/proto"

	"google.golang.org/grpc"
)

type BookServer struct {
	pb.BookServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	if err := connectDatabase(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Disconnect(context.Background())
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &BookServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
