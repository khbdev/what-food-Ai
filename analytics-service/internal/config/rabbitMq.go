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
	exchange := os.Getenv("RABBITMQ_EXCHANGE")
	routingKey := os.Getenv("S_ROUTING_KEY")
	retryKey := os.Getenv("AUTH_RETRY_ROUTING_KEY")

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

	// setup
	if err := setup(ch, exchange, routingKey, retryKey); err != nil {
		log.Fatal(err)
	}

	fmt.Println("RabbitMq Connection")

	return r
}

// =====================
// SETUP
// =====================

func setup(ch *amqp.Channel, exchange, routingKey, retryKey string) error {

	// Exchange (direct)
	if err := ch.ExchangeDeclare(
		exchange,
		"direct",
		true,  // durable
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// DLQ
	_, err := ch.QueueDeclare(
		"statik_failed_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Main queue (DLX bilan)
	args := amqp.Table{
		"x-dead-letter-exchange":    exchange,
		"x-dead-letter-routing-key": retryKey,
	}

	_, err = ch.QueueDeclare(
		"statik_success_queue",
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return err
	}

	// Bind main
	if err := ch.QueueBind(
		"statik_success_queue",
		routingKey,
		exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	// Bind DLQ
	if err := ch.QueueBind(
		"statik_failed_queue",
		"failed",
		exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}