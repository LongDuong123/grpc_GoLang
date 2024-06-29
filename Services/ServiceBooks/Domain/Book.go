package domain

import "grpc_project/Services/ServiceBooks/Application/proto"

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
	Amount int64  `json:"amount"`
}

type RespositoryBook interface {
	GetBookByID(id int64) (*Book, error)
	CreateBook(book *Book) error
	UpdateBook(bookUpdate *Book) error
	DeleteBook(id int64) error
}

type BookInteractor interface {
	GetBookByID(*proto.BookID) (*proto.Book, error)
	CreateBook(*proto.Book) error
	UpdateBook(*proto.Book) error
	DeleteBook(*proto.BookID) error
}
