package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/entity"
)

type Resolver struct {
	Repository entity.RepositoryOrdemInterface
}
