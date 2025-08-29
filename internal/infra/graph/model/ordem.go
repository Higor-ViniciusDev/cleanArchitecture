package model

// Ordem representa uma ordem de compra
type Ordem struct {
	ID    string  `json:"id"`
	Preco float64 `json:"preco"`
	Taxa  float64 `json:"taxa"`
	Valor float64 `json:"valor,omitempty"`
}
