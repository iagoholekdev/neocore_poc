package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"apicase/internal/domain/exame/service"
	"pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	// Conecta no micro-extrator
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("falha ao conectar com micro-extrator: %v", err)
	}
	defer conn.Close()

	// Inicializa client do extrator no service
	service.InitExtratorClient(conn)

	for i := 0; i < 15000; i++ {
		req := &proto.ExtractRequest{
			FilePath:   "arquivo_fake.csv",
			DomainType: proto.DomainType_EXAME,
		}
		// Chama o método do serviço ExameService
		resp, err := service.NewExameService().ExtractData(context.Background(), req)
		if err != nil {
			log.Fatalf("erro na requisição: %v", err)
		}
		fmt.Println("request numero: ", i)
		log.Printf("Resposta do micro-extrator: success=%v, message=%s", resp.Success, resp.Message)
	}

	// Para simular algum delay, se quiser
	time.Sleep(time.Second * 1000)
}
