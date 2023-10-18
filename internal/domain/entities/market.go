package entities

type Market struct {
	Name        string
	FantasyName string
	Cnpj        *CNPJ
	Address     string
}

func NewMarket(name, address string, cnpj *CNPJ) *Market {
	return &Market{
		Name:        name,
		FantasyName: "",
		Cnpj:        cnpj,
		Address:     address,
	}
}

func (m *Market) UpdateFantasyName(fantasyName string) {
	m.FantasyName = fantasyName
}

func (m *Market) UpdateName(name string) {
	m.Name = name
}

func (m *Market) UpdateAddress(address string) {
	m.Address = address
}

type Purchase struct {
	Market Market
	Items  []*Item
	URL    string
}

func NewPurchase(url string) *Purchase {
	return &Purchase{
		URL: url,
	}
}

func (p *Purchase) AddItem(item *Item) {
	p.Items = append(p.Items, item)
}

func (p *Purchase) UpdateMarket(market Market) {
	p.Market = market
}
