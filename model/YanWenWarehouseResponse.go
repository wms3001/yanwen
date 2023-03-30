package model

type YanWenWarehouseResponse struct {
	YanWenResponse
	Data []YanWenWarehouse `json:"data"`
}
