package web

import (
	"encoding/json"
	"net/http"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	"github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
)

type OrdensHandler struct {
	Repository         entity.RepositoryOrdemInterface
	EventoOrdemCreated events.EventoInterface
	EventoDisparador   events.EventoDisparadorInterface
}

func NewOrdensHandler(repository entity.RepositoryOrdemInterface, EventoOrdemCreated events.EventoInterface, EventoDisparador events.EventoDisparadorInterface) *OrdensHandler {
	return &OrdensHandler{
		Repository:         repository,
		EventoOrdemCreated: EventoOrdemCreated,
		EventoDisparador:   EventoDisparador,
	}
}

func (o *OrdensHandler) CriarOrdem(w http.ResponseWriter, r *http.Request) {
	dto := usecase.OrdemInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateOrdemUseCase(o.Repository, o.EventoOrdemCreated, o.EventoDisparador)

	outputDTO, err := useCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputDTO)
}
