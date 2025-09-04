package usecase

import (
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
)

type OrdemUseCase struct {
	repository         entity.RepositoryOrdemInterface
	eventoOrdemCreated events.EventoInterface
	eventoDisparador   events.EventoDisparadorInterface
}

func NewCreateOrdemUseCase(OrderRepository entity.RepositoryOrdemInterface, EventoOrdemCreated events.EventoInterface, DisparadorEvento events.EventoDisparadorInterface) *OrdemUseCase {
	return &OrdemUseCase{
		repository:         OrderRepository,
		eventoOrdemCreated: EventoOrdemCreated,
		eventoDisparador:   DisparadorEvento,
	}
}

type OrdemInputDTO struct {
	ID    string  `json:"id"`
	Preco float64 `json:"preco"`
	Taxa  float64 `json:"taxa"`
}

type OrdemOutputDTO struct {
	ID    string  `json:"id"`
	Preco float64 `json:"preco"`
	Taxa  float64 `json:"taxa"`
	Valor float64 `json:"valor"`
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

	dtoRetorno := &OrdemOutputDTO{
		ID:    ordem.ID,
		Preco: ordem.Preco,
		Taxa:  ordem.Taxa,
		Valor: ordem.Valor,
	}

	u.eventoOrdemCreated.SetValues(dtoRetorno)
	u.eventoDisparador.Disparador(u.eventoOrdemCreated)

	return dtoRetorno, nil
}
