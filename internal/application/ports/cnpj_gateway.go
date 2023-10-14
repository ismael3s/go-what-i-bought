package ports

import "github.com/ismael3s/gowhatibought/internal/domain/entities"

type EnterpriseByCNPJ struct {
	FantasyName string
}

type CNPJGateway interface {
	FindEnterpriseByCNPJ(cnpj *entities.CNPJ) (*EnterpriseByCNPJ, error)
}
