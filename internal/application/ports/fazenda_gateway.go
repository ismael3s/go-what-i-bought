package ports

import "github.com/ismael3s/gowhatibought/internal/domain/entities"

type FazendaGateway interface {
	FindBuyInfo(url string) (*entities.Market, error)
}
