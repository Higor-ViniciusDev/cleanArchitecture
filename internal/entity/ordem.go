package entity

import "errors"

type Ordem struct {
	ID    string
	Preco float64
	Taxa  float64
	Valor float64
}

func NovaOrdem(id string, preco, taxa float64) (*Ordem, error) {
	novarOrdem := &Ordem{
		ID:    id,
		Preco: preco,
		Taxa:  taxa,
	}

	err := novarOrdem.Validar()

	if err != nil {
		return nil, err
	}

	return novarOrdem, nil
}

func (o *Ordem) Validar() error {
	if o.ID == "" {
		return errors.New("ID não pode ser vazio")
	}
	if o.Preco <= 0 {
		return errors.New("Preço deve ser maior que zero")
	}
	if o.Taxa < 0 {
		return errors.New("Taxa não pode ser negativa")
	}
	return nil
}

func (o *Ordem) CalcularValorFinal() error {
	o.Valor = o.Preco + o.Taxa
	err := o.Validar()

	if err != nil {
		return err
	}

	return nil
}
