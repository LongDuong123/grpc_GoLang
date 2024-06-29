package grpcbook

import (
	"context"
	"grpc_project/Services/ServiceOrder/Application/proto"
	domain "grpc_project/Services/ServiceOrder/Domain"
	"io"
)

type RepositoryGrpc struct {
	grpc_book *GrpcBook
}

func NewRepositoryGprc(grpc *GrpcBook) domain.BookRepository {
	return &RepositoryGrpc{grpc_book: grpc}
}

func (repository *RepositoryGrpc) GetBookByID(id []int64) ([]domain.Book, error) {
	stream, err := repository.grpc_book.Client.ReadBook(context.Background())
	if err != nil {
		return nil, err
	}
	for _, IdBook := range id {
		proto_IdBook := &proto.BookID{Id: IdBook}
		if err := stream.Send(proto_IdBook); err != nil {
			return nil, err
		}
	}
	if err := stream.CloseSend(); err != nil {
		return nil, err
	}
	var list_book []domain.Book
	for {
		receivedBook, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		book := domain.Book{
			BookID: receivedBook.Id,
			Price:  receivedBook.Price,
			Amount: receivedBook.Amount,
		}
		list_book = append(list_book, book)
	}
	return list_book, nil
}

func (repository *RepositoryGrpc) UpdateBook(book domain.Book) (*proto.BookResponse, error) {
	proto_book := &proto.Book{Id: book.BookID, Price: book.Price, Amount: book.Amount}
	update, err := repository.grpc_book.Client.UpdateBook(context.Background(), proto_book)
	if err != nil {
		return nil, err
	}
	return update, nil
}
