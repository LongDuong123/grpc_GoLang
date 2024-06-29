package main

import (
	rabbitmq "grpc_project/Services/ServiceSendEmail/Application/RabbitMQ"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Connect AMQP fail", err)
	}
	defer conn.Close()
	err = rabbitmq.NewConsumer(conn).ConsumeMessages()
	if err != nil {
		log.Fatal("Fail connect : ", err)
	}
}
