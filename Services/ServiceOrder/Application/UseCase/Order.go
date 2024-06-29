package usecase

import (
	"errors"
	rabbitmqUseCase "grpc_project/Services/ServiceOrder/Application/RabbitMQ"
	domain "grpc_project/Services/ServiceOrder/Domain"
)

type OrderInteractor struct {
	BookRepository domain.BookRepository
	Email          rabbitmqUseCase.SendEmail
}

func NewOrderInteractor(br domain.BookRepository, rb rabbitmqUseCase.SendEmail) domain.OrderInteractor {
	return &OrderInteractor{BookRepository: br, Email: rb}
}

func (orderUseCase *OrderInteractor) CreateOrder(order *domain.Order) (*domain.Order, error) {
	books, err := orderUseCase.BookRepository.GetBookByID(order.BookID)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(books); i++ {
		if order.AmountRequired[i] > books[i].Amount {
			return nil, errors.New("not enough quantity")
		}
		books[i].Amount -= order.AmountRequired[i]
		order.Total += (books[i].Price * order.AmountRequired[i])
	}
	go func() error {
		for i := 0; i < len(books); i++ {
			response, err := orderUseCase.BookRepository.UpdateBook(books[i])
			if err != nil || response == nil {
				return err
			}
		}
		return nil
	}()
	err = orderUseCase.Email.Publish([]byte("Order successful"))
	if err != nil {
		return nil, err
	}
	return order, nil
}
