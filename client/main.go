package main

import (
	"context"
	"log"

	pb "grpc_project/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	stream, err := client.CreateBook(context.Background())
	if err != nil {
		log.Fatalf("CreateBook call failed: %v", err)
	}
	books := []*pb.Book{
		{Id: "1", Title: "Book 1", Author: "Author 1"},
		{Id: "2", Title: "Book 2", Author: "Author 2"},
	}
	for _, book := range books {
		if err := stream.Send(book); err != nil {
			log.Fatalf("Send failed: %v", err)
		}
		log.Printf("Book sent: %v", book)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Failed to close stream: %v", err)
	}

	response, err := stream.Recv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	log.Printf("CreateBook response: %v", response)
	bookID := &pb.BookID{Id: 1}
	readResponse, err := client.ReadBook(context.Background(), bookID)
	if err != nil {
		log.Fatalf("ReadBook call failed: %v", err)
	}
	log.Printf("ReadBook response: %v", readResponse)

	// Gọi phương thức UpdateBook
	updateResponse, err := client.UpdateBook(context.Background(), books[0])
	if err != nil {
		log.Fatalf("UpdateBook call failed: %v", err)
	}
	log.Printf("UpdateBook response: %v", updateResponse)

	// Gọi phương thức DeleteBook
	deleteResponse, err := client.DeleteBook(context.Background(), bookID)
	if err != nil {
		log.Fatalf("DeleteBook call failed: %v", err)
	}
	log.Printf("DeleteBook response: %v", deleteResponse)
}
