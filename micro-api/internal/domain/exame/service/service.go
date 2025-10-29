package service

import (
	"context"
	"pkg/proto"

	"google.golang.org/grpc"
)

// ExameService implementa o gRPC server do ExtractorService
type ExameService struct {
	proto.UnimplementedExtractorServiceServer
	// Aqui você pode colocar dependências, DB, etc.
}

var extratorClient proto.ExtractorServiceClient

func InitExtratorClient(conn *grpc.ClientConn) {
	extratorClient = proto.NewExtractorServiceClient(conn)
}

// NewExameService cria uma instância do serviço
func NewExameService() *ExameService {
	return &ExameService{}
}

// ExtractData implementa o método gRPC
func (s *ExameService) ExtractData(ctx context.Context, req *proto.ExtractRequest) (*proto.ExtractResponse, error) {
	if req.FilePath == "" {
		return &proto.ExtractResponse{
			Success: false,
			Message: "file_path vazio",
		}, nil
	}

	// Aqui você chamaria o microserviço extrator
	resp, err := extratorClient.ExtractData(ctx, req)
	if err != nil {
		return &proto.ExtractResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &proto.ExtractResponse{
		Success: resp.Success,
		Message: resp.Message,
	}, nil
}
