package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrdem(t *testing.T) {
	order, err := NovaOrdem("teste", 100.50, 10)

	assert.Nil(t, err, "Não pode haver error na criacao")
	assert.NotEmpty(t, order.ID, "ID não pode ser vazio")
	assert.NotEmpty(t, order.Preco, "Preço não pode ser vazio")
	assert.NotEmpty(t, order.Taxa, "Taxa não pode ser vazio")
}

func TestValidarOrdem(t *testing.T) {
	order, err := NovaOrdem("teste", 100.50, 10)
	assert.Nil(t, err, "Não pode haver error na criacao")

	err = order.Validar()
	assert.Nil(t, err, "Deve ser valido")

	order.ID = ""
	err = order.Validar()
	assert.Error(t, err, "ID não pode ser vazio")

	order.ID = "teste"
	order.Preco = 0
	err = order.Validar()
	assert.Error(t, err, "Preço deve ser maior que zero")

	order.Preco = 100.50
	order.Taxa = -1
	err = order.Validar()
	assert.Error(t, err, "Taxa não pode ser negativa")
}

func TestCalculaValorOrdem(t *testing.T) {
	order, err := NovaOrdem("teste", 100.50, 10)
	assert.Nil(t, err, "Não pode haver error na criacao")

	valor := order.CalcularValorFinal()
	assert.Equal(t, 110.50, valor, "Valor deve ser igual a soma do preço com a taxa")
}
