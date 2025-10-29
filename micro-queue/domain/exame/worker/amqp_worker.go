package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"queuetest/domain/exame/model"
	"queuetest/domain/exame/service"
	"queuetest/pkg/queue"
)

type ExameWorker struct {
	rabbit  *queue.RabbitMQ
	service *service.ExameService
}

func NewExameWorker(rabbit *queue.RabbitMQ, service *service.ExameService) *ExameWorker {
	return &ExameWorker{rabbit: rabbit, service: service}
}

func (w *ExameWorker) StartConsumers() {
	err := w.rabbit.Consume("exame.consumidor", w.handleExameCriado)
	queue.FailOnError(err, "Erro ao registrar consumer exame_criado")

	log.Println("[Worker] Consumer exame_criado iniciado")
}

// handler de evento exame_criado
func (w *ExameWorker) handleExameCriado(body []byte) {
	var exame model.Exame
	if err := json.Unmarshal(body, &exame); err != nil {
		log.Printf("[Worker] Erro ao decodificar exame_criado: %v", err)
		return
	}

	fmt.Println(exame)

	log.Printf("[Worker] Exame processado com sucesso: %+v", exame)
}
