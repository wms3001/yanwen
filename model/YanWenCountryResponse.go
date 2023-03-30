package model

type YanWenCountryResponse struct {
	YanWenResponse
	Data []YanWenCountry `json:"data"`
}
