// domain/exame/service/service.go
package service

import (
	"encoding/json"
	"log"
	"queuetest/pkg/queue"
)

type ExameService struct {
	rabbit *queue.RabbitMQ
}

func NewExameService(rabbit *queue.RabbitMQ) *ExameService {
	return &ExameService{rabbit: rabbit}
}

type Exame struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *ExameService) CriarExame(exame Exame) error {
	body, _ := json.Marshal(exame)

	err := s.rabbit.Publish("", "exame.consumidor", body)
	if err != nil {
		return err
	}

	log.Printf("Exame criado e publicado: %+v", exame)
	return nil
}
