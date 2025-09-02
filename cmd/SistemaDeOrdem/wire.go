package main

import (
	"database/sql"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/events"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/database"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/usecase"
	pkgEvento "github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrdemRepository,
	wire.Bind(new(entity.RepositoryOrdemInterface), new(*database.OrdemRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	pkgEvento.NewEventoDisparador,
	events.NewOrdemCreated,
	wire.Bind(new(pkgEvento.EventoInterface), new(*events.OrdemCreated)),
	wire.Bind(new(pkgEvento.EventoDisparadorInterface), new(*pkgEvento.EventoDisparador)),
)

var setOrderCreatedEvent = wire.NewSet(
	events.NewOrdemCreated,
	wire.Bind(new(pkgEvento.EventoInterface), new(*events.OrdemCreated)),
)

func NewCreateOrdemUseCase(db *sql.DB, disparador *pkgEvento.EventoDisparador) *usecase.OrdemUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrdemUseCase,
	)

	return &usecase.OrdemUseCase{}
}

// func NewCreateHandlerWeb(db *sql.DB, disparador *pkgEvento.EventoDisparador) *web.OrdensHandler {
// 	wire.B
// }
