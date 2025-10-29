package main

import (
	"log"
	"queuetest/domain/exame/service"
	"queuetest/pkg/queue"
)

func main() {
	// 1️⃣ Inicializa conexão RabbitMQ (infraestrutura)
	rabbit, err := queue.NewRabbitMQ("amqp://admin:admin123@localhost:5673/")
	if err != nil {
		log.Fatalf("Erro ao conectar no RabbitMQ: %v", err)
	}
	defer rabbit.Close()

	_, err = rabbit.DeclareQueue("exame.consumidor")

	if err != nil {
		panic(err)
	}

	// 2️⃣ Inicializa o serviço de domínio (regras de negócio)
	exameService := service.NewExameService(rabbit)

	// 4️⃣ Exemplo opcional de publicação (producer)
	_ = exameService.CriarExame(service.Exame{
		ID:   "001",
		Name: "Hemograma Completo",
	})

	select {}
}
