package web

import (
	"encoding/json"
	"net/http"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
)

type OrdensHandler struct {
	Repository entity.RepositoryOrdemInterface
}

func NewOrdensHandler(repository entity.RepositoryOrdemInterface) *OrdensHandler {
	return &OrdensHandler{
		Repository: repository,
	}
}

func (o *OrdensHandler) CriarOrdem(w http.ResponseWriter, r *http.Request) {
	dto := usecase.OrdemInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateOrdemUseCase(o.Repository)

	outputDTO, err := useCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputDTO)
}
