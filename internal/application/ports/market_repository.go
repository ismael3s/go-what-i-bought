package ports

import "github.com/ismael3s/gowhatibought/internal/domain/entities"

type MarketRepository interface {
	Save(market *entities.Market) error
	FindByNFURL(url string) (*entities.Market, error)
}
