package main

import (
	"database/sql"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/database"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ordens")

	if err != nil {
		panic(err)
	}

	novoRepository := database.NewOrdemRepository(db)
	useCase := usecase.NewCreateOrdemUseCase(novoRepository)

	dto := &usecase.OrdemInputDTO{
		ID:    "teste",
		Preco: 120.2,
		Taxa:  5.5,
	}

	useCase.Execute(*dto)
}
