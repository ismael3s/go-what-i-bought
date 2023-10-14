package entities

import (
	"strconv"
	"strings"

	"github.com/ismael3s/gowhatibought/internal/domain/shared"
)

type Item struct {
	Name       string
	Price      float64
	Quantity   float64
	Unit       string
	TotalValue float64
}

func NewItem(name string, price string, quantity string, unit string, totalValue string) (*Item, error) {
	floatQuantity, err := extractNumberAsFloat(quantity)
	if err != nil {
		return nil, err
	}
	floatPrice, err := extractNumberAsFloat(price)
	if err != nil {
		return nil, err
	}
	floatTotalValue, err := extractNumberAsFloat(totalValue)
	if err != nil {
		return nil, err
	}
	rawUnit := strings.ReplaceAll(unit, "UN: ", "")
	return &Item{
		Name:       name,
		Price:      floatPrice,
		Quantity:   floatQuantity,
		Unit:       rawUnit,
		TotalValue: floatTotalValue,
	}, nil
}

func extractNumberAsFloat(s string) (float64, error) {
	stringWithDot := strings.Replace(s, ",", ".", 1)
	return strconv.ParseFloat(shared.ONLY_NUMBERS_REGEX.FindString(stringWithDot), 64)
}
