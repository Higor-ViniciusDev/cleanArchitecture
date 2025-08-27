package usecase

import "github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"

type OrdemUseCase struct {
	repository entity.RepositoryOrdemInterface
}

func NewCreateOrdemUseCase(OrderRepository entity.RepositoryOrdemInterface) *OrdemUseCase {
	return &OrdemUseCase{
		repository: OrderRepository,
	}
}

type OrdemInputDTO struct {
	ID    string
	Preco float64
	Taxa  float64
}

type OrdemOutputDTO struct {
	ID    string
	Preco float64
	Taxa  float64
	Valor float64
}

func (u *OrdemUseCase) Execute(input OrdemInputDTO) (*OrdemOutputDTO, error) {
	ordem, err := entity.NovaOrdem(input.ID, input.Preco, input.Taxa)
	if err != nil {
		return nil, err
	}

	err = ordem.CalcularValorFinal()
	if err != nil {
		return nil, err
	}

	err = u.repository.Salvar(ordem)
	if err != nil {
		return nil, err
	}

	return &OrdemOutputDTO{
		ID:    ordem.ID,
		Preco: ordem.Preco,
		Taxa:  ordem.Taxa,
		Valor: ordem.Valor,
	}, nil
}
