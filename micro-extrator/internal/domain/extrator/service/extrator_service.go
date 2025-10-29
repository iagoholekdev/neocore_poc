package service

import (
	"context"
	"fmt"

	"exctrator/internal/domain/extrator/model"
	"pkg/proto"
)

type ExameService struct {
	proto.UnimplementedExtractorServiceServer
	useCase *model.ExameUseCase
}

func NewExameService() *ExameService {
	return &ExameService{
		useCase: model.NewExameUseCase(),
	}
}

func (s *ExameService) ExtractData(ctx context.Context, req *proto.ExtractRequest) (*proto.ExtractResponse, error) {
	if req.FilePath == "" {
		return &proto.ExtractResponse{
			Success: false,
			Message: "file_path vazio",
		}, nil
	}

	exames, err := s.useCase.Process(req.FilePath)
	if err != nil {
		return &proto.ExtractResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	fmt.Printf("Processados %d exames\n", len(exames))

	return &proto.ExtractResponse{
		Success: true,
		Message: fmt.Sprintf("Processados %d exames com sucesso!", len(exames)),
	}, nil
}
