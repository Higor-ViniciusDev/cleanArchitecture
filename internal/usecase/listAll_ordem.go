package usecase

import (
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
)

type ListOrdemUseCase struct {
	ordemRepository entity.RepositoryOrdemInterface
}

func NewListOrdemUseCase(ordemRepo entity.RepositoryOrdemInterface) *ListOrdemUseCase {
	return &ListOrdemUseCase{
		ordemRepository: ordemRepo,
	}
}

func (uc *ListOrdemUseCase) Execute() ([]OrdemOutputDTO, error) {
	ordens, err := uc.ordemRepository.BuscarTodas()
	if err != nil {
		return nil, err
	}

	var ordensOutput []OrdemOutputDTO
	for _, ordem := range ordens {
		ordemOutput := OrdemOutputDTO{
			ID:    ordem.ID,
			Preco: ordem.Preco,
			Taxa:  ordem.Taxa,
			Valor: ordem.Valor,
		}
		ordensOutput = append(ordensOutput, ordemOutput)
	}

	return ordensOutput, nil
}
