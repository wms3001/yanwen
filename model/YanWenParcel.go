package model

type YanWenParcel struct {
	HasBattery    int             `json:"hasBattery"`
	Currency      string          `json:"currency"`
	TotalPrice    float64         `json:"totalPrice"`
	TotalQuantity int             `json:"totalQuantity"`
	TotalWeight   int             `json:"totalWeight"`
	Height        int             `json:"height"`
	Width         int             `json:"width"`
	Length        int             `json:"length"`
	Ioss          string          `json:"ioss"`
	ProductList   []YanWenProduct `json:"productList"`
}
