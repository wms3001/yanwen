package model

type YanWenOrderRe struct {
	WaybillNumber   string `json:"waybillNumber"`
	OrderNumber     string `json:"orderNumber"`
	ReferenceNumber string `json:"referenceNumber"`
	YanwenNumber    string `json:"yanwenNumber"`
}
