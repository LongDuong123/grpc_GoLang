package main

import (
	"context"
	pb "grpc_project/proto"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *BookServer) CreateBook(stream pb.BookService_CreateBookServer) error {
	for {
		Book, err := stream.Recv()
		if err != nil {
			return err
		}

		_, err = collection.InsertOne(context.Background(), Book)
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.BookResponse{Status: "Create Successful"}); err != nil {
			return err
		}
	}
}

func (s *BookServer) ReadBook(ctx context.Context, bookID *pb.BookID) (*pb.Book, error) {
	var book pb.Book

	filter := bson.M{"id": bookID.Id}

	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookServer) UpdateBook(ctx context.Context, book *pb.Book) (*pb.BookResponse, error) {
	filter := bson.M{"id": book.Id}
	updateBook := bson.M{"$set": book}

	_, err := collection.UpdateOne(context.Background(), filter, updateBook)

	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{Status: "Update Successful"}, nil
}

func (s *BookServer) DeleteBook(ctx context.Context, bookID *pb.BookID) (*pb.BookResponse, error) {
	filter := bson.M{"id": bookID}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{Status: "Delete Successful"}, nil
}
