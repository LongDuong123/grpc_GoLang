package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbiitMQConsumer struct {
	Conn *amqp.Connection
}

func NewConsumer(connect *amqp.Connection) *RabbiitMQConsumer {
	return &RabbiitMQConsumer{Conn: connect}
}

func (consumer *RabbiitMQConsumer) ConsumeMessages() error {
	ch, err := consumer.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
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
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	var forever chan struct{}
	for d := range msgs {
		log.Printf("Send email : %s", d.Body)
	}
	<-forever
	return nil
}
