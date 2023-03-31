package model

type YanWenLabelResponse struct {
	YanWenResponse
	Data YanWenLabelData `json:"data"`
}
