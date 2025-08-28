package services

import (
	"context"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdemService struct {
	pb.UnimplementedOrdemServiceServer
	repository entity.RepositoryOrdemInterface
}

func (c *OrdemService) CriarOrdem(ctx context.Context, in *pb.CriarOrdemRequest) (*pb.OrdemOutput, error) {
	orderusecase := usecase.NewCreateOrdemUseCase(c.repository)

	dto := usecase.OrdemInputDTO{
		ID:    in.Ordem.Id,
		Preco: float64(in.Ordem.Preco),
		Taxa:  float64(in.Ordem.Taxa),
	}

	output, err := orderusecase.Execute(dto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Falhar ao criar ordem: %v", err)
	}

	return &pb.OrdemOutput{
		Id:    output.ID,
		Preco: float32(output.Preco),
		Taxa:  float32(output.Taxa),
		Valor: float32(output.Valor),
	}, nil
}

func NewOrdemService(repository entity.RepositoryOrdemInterface) *OrdemService {
	return &OrdemService{
		repository: repository,
	}
}
