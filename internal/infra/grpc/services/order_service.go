package services

import (
	"context"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	pb.UnimplementedOrdemServiceServer
	CreateOrderUseCase usecase.OrdemUseCase
	ListAllUseCase     usecase.ListOrdemUseCase
}

func (c *OrderService) CriarOrdem(ctx context.Context, in *pb.CriarOrdemRequest) (*pb.OrdemOutput, error) {
	dto := usecase.OrdemInputDTO{
		ID:    in.Ordem.Id,
		Preco: float64(in.Ordem.Preco),
		Taxa:  float64(in.Ordem.Taxa),
	}

	output, err := c.CreateOrderUseCase.Execute(dto)
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

func (c *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.ListarOrdensResponse, error) {
	output, err := c.ListAllUseCase.Execute()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Falhar ao listar ordens: %v", err)
	}

	var ordens []*pb.OrdemOutput
	for _, ordem := range output {
		ordens = append(ordens, &pb.OrdemOutput{
			Id:    ordem.ID,
			Preco: float32(ordem.Preco),
			Taxa:  float32(ordem.Taxa),
			Valor: float32(ordem.Valor),
		})
	}

	return &pb.ListarOrdensResponse{Ordens: ordens}, nil
}

func NewOrderService(createOrderUseCase usecase.OrdemUseCase, listAllOrdeUseCase usecase.ListOrdemUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListAllUseCase:     listAllOrdeUseCase,
	}
}
