package usecase

import (
	"grpc_project/Services/ServiceBooks/Application/proto"
	domain "grpc_project/Services/ServiceBooks/Domain"
)

type BookInteractor struct {
	bookRepository domain.RespositoryBook
}

func NewBookInteractor(br domain.RespositoryBook) domain.BookInteractor {
	return &BookInteractor{bookRepository: br}
}

func (bookInteractor *BookInteractor) GetBookByID(id *proto.BookID) (*proto.Book, error) {
	book, err := bookInteractor.bookRepository.GetBookByID(id.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Book{Id: book.Id, Title: book.Title, Author: book.Author, Price: book.Price, Amount: book.Amount}, nil
}

func (bookInteractor *BookInteractor) CreateBook(book *proto.Book) error {
	addBook := &domain.Book{Id: book.Id, Title: book.Title, Author: book.Author, Price: book.Price, Amount: book.Amount}
	err := bookInteractor.bookRepository.CreateBook(addBook)
	if err != nil {
		return err
	}
	return nil
}

func (bookInteractor *BookInteractor) UpdateBook(book *proto.Book) error {
	updateBook := &domain.Book{Id: book.Id, Title: book.Title, Author: book.Author, Price: book.Price, Amount: book.Amount}
	err := bookInteractor.bookRepository.UpdateBook(updateBook)
	if err != nil {
		return err
	}
	return nil
}

func (bookInteractor *BookInteractor) DeleteBook(id *proto.BookID) error {
	err := bookInteractor.bookRepository.DeleteBook(id.Id)
	if err != nil {
		return err
	}
	return nil
}
