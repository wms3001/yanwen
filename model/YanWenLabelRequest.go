package model

type YanWenLabelRequest struct {
	WaybillNumber string `json:"waybillNumber"`
	PrintRemark   int    `json:"printRemark"`
}
