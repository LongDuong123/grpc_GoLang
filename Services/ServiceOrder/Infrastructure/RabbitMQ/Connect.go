package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQEmail struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func ConnectServiceEmail() (*RabbitMQEmail, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQEmail{Conn: conn, Channel: ch}, nil
}
