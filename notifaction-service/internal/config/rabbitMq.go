package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// =====================
// INIT (ONLY CONNECTION)
// =====================

func NewRabbit() *Rabbit {
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ channel error:", err)
	}

	// optional QoS (xohlasang qoldirasan)
	if err := ch.Qos(10, 0, false); err != nil {
		log.Fatal("RabbitMQ QoS error:", err)
	}

	log.Println("RabbitMQ connected")

	return &Rabbit{
		Conn:    conn,
		Channel: ch,
	}
}