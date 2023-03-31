package model

type YanWenProduct struct {
	GoodsNameCh string  `json:"goodsNameCh"`
	GoodsNameEn string  `json:"goodsNameEn"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Weight      int     `json:"weight"`
	Hscode      string  `json:"hscode"`
	Url         string  `json:"url"`
	Material    string  `json:"material"`
}
