package web

import (
	"encoding/json"
	"net/http"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/presenters"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
)

const (
	acceptJson = "application/json"
	acceptXml  = "application/xml"
)

type OrdensHandler struct {
	CreateUseCase  *usecase.OrdemUseCase
	ListAllUseCase *usecase.ListOrdemUseCase
}

func NewOrdensHandler(createUseCase *usecase.OrdemUseCase, listAllUseCase *usecase.ListOrdemUseCase) *OrdensHandler {
	return &OrdensHandler{
		CreateUseCase:  createUseCase,
		ListAllUseCase: listAllUseCase,
	}
}

func (o *OrdensHandler) CriarOrdem(w http.ResponseWriter, r *http.Request) {
	dto := usecase.OrdemInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outputDTO, err := o.CreateUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputDTO)
}

func (h *OrdensHandler) ListarOrdens(w http.ResponseWriter, r *http.Request) {
	outputDTO, err := h.ListAllUseCase.Execute(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accept := r.Header.Get("Accept")
	switch accept {
	case acceptXml:
		presenter := presenters.NewOrderPresenter()
		retorno := presenter.ToXML(&outputDTO)
		writeResponse(w, http.StatusOK, acceptXml, retorno)
	default:
		writeJSON(w, http.StatusOK, outputDTO)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", acceptJson)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Erro ao serializar dados", http.StatusInternalServerError)
	}
}

func writeResponse(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
