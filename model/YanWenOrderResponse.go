package model

type YanWenOrderResponse struct {
	YanWenResponse
	Data YanWenOrderRe `json:"data"`
}
