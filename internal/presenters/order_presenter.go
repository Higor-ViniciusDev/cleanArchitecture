package presenters

import (
	"encoding/json"
	"encoding/xml"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
)

type OrderPresenter struct{}

// Convert um array de OrderOutputDTO para JSON e retonar no formato string
func (o *OrderPresenter) ToJSON(dtosResult *[]usecase.OrdemOutputDTO) string {
	jsonData, _ := json.Marshal(dtosResult)
	return string(jsonData)
}

func NewOrderPresenter() *OrderPresenter {
	return &OrderPresenter{}
}

func (o *OrderPresenter) ToXML(dto *[]usecase.OrdemOutputDTO) []byte {
	xmlData, _ := xml.Marshal(dto)
	return xmlData
}
