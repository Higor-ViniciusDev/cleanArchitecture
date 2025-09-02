package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
	"github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
)

type Resolver struct {
	Repository         entity.RepositoryOrdemInterface
	EventoOrdemCreated events.EventoInterface
	EventoDisparador   events.EventoDisparadorInterface
}
