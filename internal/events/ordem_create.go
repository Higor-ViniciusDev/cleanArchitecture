package events

import "time"

type OrdemCreated struct {
	Nome   string
	Values interface{}
}

func NewOrdemCreated() *OrdemCreated {
	return &OrdemCreated{
		Nome: "OrdemCreated",
	}
}

func (e *OrdemCreated) GetNome() string {
	return e.Nome
}

func (e *OrdemCreated) GetValues() interface{} {
	return e.Values
}

func (e *OrdemCreated) SetValues(values interface{}) {
	e.Values = values
}

func (e *OrdemCreated) GetDateTime() time.Time {
	return time.Now()
}
