package model

type YanWenOrderInfoResponse struct {
	YanWenResponse
	Data YanWenOrderInfoData `json:"data"`
}
