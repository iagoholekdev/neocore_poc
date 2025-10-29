package main

import (
	"log"
	"queueconsumer/internal/domain/exame/worker"
	"queueconsumer/pkg/queue"
)

func main() {
	rabbit, err := queue.NewRabbitMQ("amqp://admin:admin123@localhost:5673/")
	queue.FailOnError(err, "Erro ao conectar no RabbitMQ")
	defer rabbit.Close()

	exameWorker := worker.NewExameWorker(rabbit)

	// inicia consumers
	exameWorker.StartConsumers()

	log.Println("[Consumer] Worker de exames iniciado...")
	select {}
}
