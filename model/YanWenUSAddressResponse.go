package model

type YanWenUSAddressResponse struct {
	YanWenResponse
	Data YanWenUSAddressData `json:"data"`
}
