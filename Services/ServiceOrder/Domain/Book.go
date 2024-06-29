package domain

import "grpc_project/Services/ServiceOrder/Application/proto"

type Book struct {
	BookID int64
	Price  int64
	Amount int64
}

type BookRepository interface {
	GetBookByID(id []int64) ([]Book, error)
	UpdateBook(Book) (*proto.BookResponse, error)
}
