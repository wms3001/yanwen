package model

type YanWenCancelRequest struct {
	WaybillNumber string `json:"waybillNumber"`
	Note          string `json:"note"`
}
