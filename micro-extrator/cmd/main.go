package main

import (
	"log"
	"net"

	"exctrator/internal/domain/extrator/service"
	"pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	// Cria servidor gRPC
	grpcServer := grpc.NewServer()

	// Cria serviço do extrator
	extratorService := service.NewExameService() // ou NewExtratorService() se você quiser outro nome
	proto.RegisterExtractorServiceServer(grpcServer, extratorService)

	// Escuta porta 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("falha ao escutar porta: %v", err)
	}
	log.Println("micro-extrator rodando na porta 50051")

	// Serve gRPC
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("falha no server gRPC: %v", err)
	}
}
