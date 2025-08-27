package database

import (
	"database/sql"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
)

type OrdemRepository struct {
	db *sql.DB
}

func NewOrdemRepository(db *sql.DB) *OrdemRepository {
	return &OrdemRepository{
		db: db,
	}
}

func (r *OrdemRepository) Salvar(ordem *entity.Ordem) error {
	_, err := r.db.Exec("INSERT INTO ordens (id, preco, taxa, valor) VALUES (?, ?, ?, ?)",
		ordem.ID, ordem.Preco, ordem.Taxa, ordem.Valor)
	return err
}

func (r *OrdemRepository) BuscarPorID(id string) (*entity.Ordem, error) {
	row := r.db.QueryRow("SELECT id, preco, taxa, valor FROM ordens WHERE id = ?", id)
	var ordem entity.Ordem
	err := row.Scan(&ordem.ID, &ordem.Preco, &ordem.Taxa, &ordem.Valor)
	if err != nil {
		return nil, err
	}
	return &ordem, nil
}

func (r *OrdemRepository) BuscarTodas() ([]*entity.Ordem, error) {
	rows, err := r.db.Query("SELECT id, preco, taxa, valor FROM ordens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []*entity.Ordem
	for rows.Next() {
		var ordem entity.Ordem
		err := rows.Scan(&ordem.ID, &ordem.Preco, &ordem.Taxa, &ordem.Valor)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, &ordem)
	}
	return ordens, nil
}
