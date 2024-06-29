package grpcbook

import (
	"grpc_project/Services/ServiceOrder/Application/proto"

	"google.golang.org/grpc"
)

type GrpcBook struct {
	Conn   *grpc.ClientConn
	Client proto.BookServiceClient
}

func ConnectServiceBooks() (*GrpcBook, error) {
	connect, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := proto.NewBookServiceClient(connect)
	return &GrpcBook{Conn: connect, Client: client}, nil
}
