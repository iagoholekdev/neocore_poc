package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) DeclareQueue(name string) (amqp.Queue, error) {
	q, err := r.ch.QueueDeclare(
		name,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	return q, err
}

func (r *RabbitMQ) DeclareExchange(name, kind string) error {
	return r.ch.ExchangeDeclare(
		name,
		kind,  // fanout, direct, topic, headers
		true,  // durable
		false, // auto-delete
		false, // internal
		false, // no-wait
		nil,   // args
	)
}

func (r *RabbitMQ) BindQueue(queue, exchange, routingKey string) error {
	return r.ch.QueueBind(
		queue,
		routingKey,
		exchange,
		false,
		nil,
	)
}

func (r *RabbitMQ) Close() {
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
	log.Println("[RabbitMQ] Conexão encerrada.")
}
