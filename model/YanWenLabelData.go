package model

type YanWenLabelData struct {
	WaybillNumber string `json:"waybillNumber"`
	IsSuccess     bool   `json:"isSuccess"`
	ErrorMsg      string `json:"errorMsg"`
	Base64String  string `json:"base64String"`
	LabelType     string `json:"labelType"`
}
