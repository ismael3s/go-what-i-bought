package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ismael3s/gowhatibought/internal/application/ports"
	"github.com/ismael3s/gowhatibought/internal/domain/entities"
)

type BrasilApiGateway struct{}

func NewBrasilApiGateway() *BrasilApiGateway {
	return &BrasilApiGateway{}
}

type BrasilAPICNPJResponse struct {
	FantasyName string `json:"nome_fantasia"`
}

func (g *BrasilApiGateway) FindEnterpriseByCNPJ(cnpj *entities.CNPJ) (*ports.EnterpriseByCNPJ, error) {
	var response BrasilAPICNPJResponse
	resp, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%s", cnpj.RawValue()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &ports.EnterpriseByCNPJ{
		FantasyName: response.FantasyName,
	}, nil
}
