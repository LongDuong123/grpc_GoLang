package controllers

import (
	"context"
	"grpc_project/Services/ServiceBooks/Application/proto"
	domain "grpc_project/Services/ServiceBooks/Domain"
	"io"
)

type BookServer struct {
	BookUseCase domain.BookInteractor
	proto.BookServiceServer
}

func (s *BookServer) CreateBook(ctx context.Context, book *proto.Book) (*proto.BookResponse, error) {

	err := s.BookUseCase.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return &proto.BookResponse{Status: "Create Book Successful"}, nil
}

func (s *BookServer) ReadBook(stream proto.BookService_ReadBookServer) error {
	for {
		bookID, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		book, err := s.BookUseCase.GetBookByID(bookID)
		if err != nil {
			return err
		}
		if err = stream.Send(book); err != nil {
			return err
		}
	}
}

func (s *BookServer) UpdateBook(ctx context.Context, book *proto.Book) (*proto.BookResponse, error) {

	err := s.BookUseCase.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return &proto.BookResponse{Status: "Update Successful"}, nil
}

func (s *BookServer) DeleteBook(ctx context.Context, bookID *proto.BookID) (*proto.BookResponse, error) {

	err := s.BookUseCase.DeleteBook(bookID)
	if err != nil {
		return nil, err
	}

	return &proto.BookResponse{Status: "Delete Successful"}, nil
}
