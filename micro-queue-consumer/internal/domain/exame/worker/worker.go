package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"queueconsumer/internal/domain/exame/model"

	"queueconsumer/pkg/queue"
)

type ExameWorker struct {
	rabbit *queue.RabbitMQ
}

func NewExameWorker(rabbit *queue.RabbitMQ) *ExameWorker {
	return &ExameWorker{rabbit: rabbit}
}

func (w *ExameWorker) StartConsumers() {
	err := w.rabbit.Consume("exame.consumidor", w.handleExameCriado)
	queue.FailOnError(err, "Erro ao registrar consumer exame_criado")
}

func (w *ExameWorker) handleExameCriado(body []byte) {
	var exame model.Exame
	if err := json.Unmarshal(body, &exame); err != nil {
		log.Printf("[Worker] Erro ao decodificar mensagem: %v", err)
		return
	}

	fmt.Println(exame)

	log.Printf("[Worker] Exame processado com sucesso: %+v", exame)
}
