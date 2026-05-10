package config

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// =====================
// INIT
// =====================

func NewRabbit() *Rabbit {
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// Prefetch
	if err := ch.Qos(10, 0, false); err != nil {
		log.Fatal(err)
	}

	r := &Rabbit{
		Conn:    conn,
		Channel: ch,
	}

	fmt.Println("RabbitMq Connection")

	return r
}
