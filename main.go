package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// change to your environment
	// user guest is only allow to connect from localhost
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ server")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to declare queue")

	var body []byte
	if len(os.Args) > 1 {
		body = []byte(strings.Join(os.Args[1:], " "))
	} else {
		body = []byte("Hello World!!!")
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body})
	handleError(err, "Failed to publish a message")
}
