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

type FindNfDataUseCaseOutput struct {
	Url         string       `json:"url"`
	FantasyName string       `json:"fantasyName"`
	FullAddress string       `json:"fullAddress"`
	Cnpj        string       `json:"cnpj"`
	Items       []ItemOutput `json:"items"`
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
	market, err := u.fazendaGateway.FindBuyInfo(input.Url)
	if err != nil {
		return nil, err
	}
	enterprise, err := u.cnpjGateway.FindEnterpriseByCNPJ(market.Cnpj)
	if err != nil {
		return nil, err
	}
	market.UpdateName("")
	market.UpdateFantasyName(enterprise.FantasyName)
	output := &FindNfDataUseCaseOutput{
		Url:         input.Url,
		FantasyName: enterprise.FantasyName,
		FullAddress: market.FullAddress,
		Cnpj:        market.Cnpj.RawValue(),
		Items:       []ItemOutput{},
	}
	for _, item := range market.Items {
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
