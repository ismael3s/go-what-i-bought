package entities

import "github.com/ismael3s/gowhatibought/internal/domain/shared"

type CNPJ struct {
	FormatedValue string
}

func (c *CNPJ) RawValue() string {
	return shared.NON_NUMERIC_REGEX.ReplaceAllString(c.FormatedValue, "")
}

func NewCNPJ(rawValue string) *CNPJ {
	return &CNPJ{
		FormatedValue: shared.ONLY_NUMBERS_REGEX.FindString(rawValue),
	}
}
