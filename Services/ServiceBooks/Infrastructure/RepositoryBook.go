package infrastructure

import (
	"context"
	domain "grpc_project/Services/ServiceBooks/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RespositoryBook struct {
	database *Database
}

func NewRepositoryBook(dtb *Database) domain.RespositoryBook {
	return &RespositoryBook{database: dtb}
}

func (rsb *RespositoryBook) GetBookByID(id int64) (*domain.Book, error) {
	var book domain.Book
	filter := bson.M{"id": id}
	err := rsb.database.collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (rsb *RespositoryBook) CreateBook(book *domain.Book) error {
	addBook := bson.M{"$set": book}
	_, err := rsb.database.collection.InsertOne(context.Background(), addBook)
	if err != nil {
		return err
	}
	return nil
}

func (rsb *RespositoryBook) UpdateBook(book *domain.Book) error {
	updateBook := bson.M{}
	if book.Title != "" {
		updateBook["title"] = book.Title
	}
	if book.Author != "" {
		updateBook["author"] = book.Author
	}
	if book.Price != -1 {
		updateBook["price"] = book.Price
	}
	if book.Amount != -1 {
		updateBook["amount"] = book.Amount
	}
	updateBook = bson.M{"$set": updateBook}
	filter := bson.M{"id": book.Id}
	_, err := rsb.database.collection.UpdateOne(context.Background(), filter, updateBook, options.Update().SetUpsert(false))
	if err != nil {
		return err
	}
	return nil
}

func (rsb *RespositoryBook) DeleteBook(id int64) error {
	filter := bson.M{"id": id}
	_, err := rsb.database.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
