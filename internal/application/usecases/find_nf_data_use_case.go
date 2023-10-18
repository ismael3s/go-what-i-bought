package usecases

import (
	"github.com/ismael3s/gowhatibought/internal/application/ports"
)

type FindNfDataUseCaseInput struct {
	Url string `json:"url"`
}

type ItemOutput struct {
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
	Unit       string  `json:"unit"`
}

type findNfDataUseCaseMarketOutput struct {
	Name        string `json:"name"`
	FantasyName string `json:"fantasyName"`
	Cnpj        string `json:"cnpj"`
	Address     string `json:"address"`
}

type FindNfDataUseCaseOutput struct {
	Url    string                        `json:"url"`
	Market findNfDataUseCaseMarketOutput `json:"market"`
	Items  []ItemOutput                  `json:"items"`
}

type FindNfDataUseCase struct {
	fazendaGateway   ports.FazendaGateway
	cnpjGateway      ports.CNPJGateway
	marketRepository ports.MarketRepository
}

func NewFindNfDataUsecase(fazendaGateway ports.FazendaGateway, cnpjGateway ports.CNPJGateway) *FindNfDataUseCase {
	return &FindNfDataUseCase{
		fazendaGateway: fazendaGateway,
		cnpjGateway:    cnpjGateway,
	}
}

func (u *FindNfDataUseCase) Execute(input FindNfDataUseCaseInput) (*FindNfDataUseCaseOutput, error) {
	purchase, err := u.fazendaGateway.FindBuyInfo(input.Url)
	if err != nil {
		return nil, err
	}
	enterprise, err := u.cnpjGateway.FindEnterpriseByCNPJ(purchase.Market.Cnpj)
	if err != nil {
		return nil, err
	}
	purchase.Market.UpdateFantasyName(enterprise.FantasyName)
	output := &FindNfDataUseCaseOutput{
		Url: input.Url,
		Market: findNfDataUseCaseMarketOutput{
			Name:        purchase.Market.Name,
			FantasyName: purchase.Market.FantasyName,
			Cnpj:        purchase.Market.Cnpj.RawValue(),
			Address:     purchase.Market.Address,
		},
		Items: []ItemOutput{},
	}
	for _, item := range purchase.Items {
		output.Items = append(output.Items, ItemOutput{
			Name:       item.Name,
			Quantity:   item.Quantity,
			UnitPrice:  item.Price,
			TotalPrice: item.TotalValue,
			Unit:       item.Unit,
		})
	}
	return output, nil
}
