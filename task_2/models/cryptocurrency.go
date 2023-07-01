package models

type Сryptocurrency struct {
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	CurPrice float64 `json:"current_price"`
}
