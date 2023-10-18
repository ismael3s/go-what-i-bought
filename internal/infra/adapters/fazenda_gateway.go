package adapters

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/ismael3s/gowhatibought/internal/domain/entities"
)

type FazendaGateway struct{}

func NewFazendaGateway() *FazendaGateway {
	return &FazendaGateway{}
}

func (f *FazendaGateway) FindBuyInfo(url string) (*entities.Purchase, error) {
	collector := colly.NewCollector()
	collector.CacheDir = "./tmp"
	var lastErr error
	purchase := entities.NewPurchase(url)
	collector.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		name := e.ChildText("td:nth-child(1) .txtTit")
		quantity := e.ChildText("td:nth-child(1) .Rqtd")
		price := e.ChildText("td:nth-child(1) .RvlUnit")
		unit := e.ChildText("td:nth-child(1) .RUN")
		totalValue := e.ChildText("td:nth-child(2) span.valor")
		item, err := entities.NewItem(name, price, quantity, unit, totalValue)
		if err != nil {
			lastErr = err
			return
		}
		purchase.AddItem(item)
	})
	collector.OnHTML("div#conteudo .txtCenter", func(e *colly.HTMLElement) {
		cnpj := entities.NewCNPJ(e.ChildText(".text:nth-child(2)"))
		address := strings.ReplaceAll(
			strings.ReplaceAll(strings.TrimSpace(e.ChildText(".text:nth-child(3)")), "\n", ""),
			"\t\t", " ")
		market := entities.NewMarket("", address, cnpj)
		purchase.UpdateMarket(*market)
	})
	collector.Visit(url)
	return purchase, lastErr
}
