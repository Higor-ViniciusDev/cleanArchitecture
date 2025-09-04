//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/events"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/database"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	pkgEvento "github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrdemRepository,
	wire.Bind(new(entity.RepositoryOrdemInterface), new(*database.OrdemRepository)),
)

var setOrderCreatedEvent = wire.NewSet(
	events.NewOrdemCreated,
	wire.Bind(new(pkgEvento.EventoInterface), new(*events.OrdemCreated)),
)

var setDisparadorEvento = wire.NewSet(
	pkgEvento.NewEventoDisparador,
	events.NewOrdemCreated,
	wire.Bind(new(pkgEvento.EventoDisparadorInterface), new(*pkgEvento.EventoDisparador)),
	wire.Bind(new(pkgEvento.EventoInterface), new(*events.OrdemCreated)),
)

func NewCreateOrdemUseCaseInje(db *sql.DB, disparador pkgEvento.EventoDisparadorInterface) *usecase.OrdemUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrdemUseCase,
	)

	return &usecase.OrdemUseCase{}
}

func NewListAllOrdemUseCaseInje(db *sql.DB) *usecase.ListOrdemUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdemUseCase,
	)

	return &usecase.ListOrdemUseCase{}
}

func NewWebOrdersHandleInje(createUseCase *usecase.OrdemUseCase, listAllUseCase *usecase.ListOrdemUseCase) *web.OrdensHandler {
	wire.Build(
		web.NewOrdensHandler,
	)

	return &web.OrdensHandler{}
}
