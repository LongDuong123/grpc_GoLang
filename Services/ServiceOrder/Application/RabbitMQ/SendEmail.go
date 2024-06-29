package rabbitmqUseCase

import (
	rabbitmq "grpc_project/Services/ServiceOrder/Infrastructure/RabbitMQ"

	"github.com/rabbitmq/amqp091-go"
)

type SendEmail struct {
	RabbitMQEmail *rabbitmq.RabbitMQEmail
}

func NewSendEmail(rb *rabbitmq.RabbitMQEmail) *SendEmail {
	return &SendEmail{RabbitMQEmail: rb}
}

func (se *SendEmail) Publish(body []byte) error {
	q, err := se.RabbitMQEmail.Channel.QueueDeclare(
		"Email",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = se.RabbitMQEmail.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
