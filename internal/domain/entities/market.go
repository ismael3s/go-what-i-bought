package entities

type Market struct {
	Name        string
	FantasyName string
	Cnpj        *CNPJ
	FullAddress string
	Items       []*Item
}

func NewMarket() *Market {
	return &Market{
		Items: []*Item{},
	}
}

func (m *Market) AddItem(item *Item) {
	m.Items = append(m.Items, item)
}

func (m *Market) UpdateFantasyName(fantasyName string) {
	m.FantasyName = fantasyName
}

func (m *Market) UpdateName(name string) {
	m.Name = name
}

func (m *Market) UpdateAddress(address string) {
	m.FullAddress = address
}
