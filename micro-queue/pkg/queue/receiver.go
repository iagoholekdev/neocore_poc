package queue

import (
	"log"
)

func (r *RabbitMQ) Consume(queue string, handler func([]byte)) error {
	msgs, err := r.ch.Consume(
		queue,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()

	log.Printf("[RabbitMQ] Consumindo da fila: %s", queue)
	return nil
}
