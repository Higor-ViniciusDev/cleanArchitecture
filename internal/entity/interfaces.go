package entity

type RepositoryOrdemInterface interface {
	Salvar(ordem *Ordem) error
	BuscarPorID(id string) (*Ordem, error)
	BuscarTodas() ([]*Ordem, error)
}
