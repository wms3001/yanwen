package model

type YanWenAddress struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	HouseNumber string `json:"houseNumber"`
	Address     string `json:"address"`
	TaxNumber   string `json:"taxNumber"`
}
