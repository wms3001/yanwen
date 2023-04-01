package model

type YanWenOrderInfoData struct {
	WaybillNumber   string `json:"waybillNumber"`
	OrderNumber     string `json:"orderNumber"`
	ReferenceNumber string `json:"referenceNumber"`
	YanwenNumber    string `json:"yanwenNumber"`
	UserId          string `json:"userId"`
	OrderSource     string `json:"orderSource"`
	ChannelId       string `json:"channelId"`
	ChannelName     string `json:"channelName"`
	DateOfReceipt   string `json:"dateOfReceipt"`
	CreateTime      string `json:"createTime"`
	CompanyCode     string `json:"companyCode"`
	IsPrint         string `json:"isPrint"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
	ReceiverInfo    struct {
		Name        string `json:"name"`
		Company     string `json:"company"`
		CountryId   string `json:"countryId"`
		CountryName string `json:"countryName"`
		Phone       string `json:"phone"`
		State       string `json:"state"`
		City        string `json:"city"`
		Email       string `json:"email"`
		ZipCode     string `json:"zipCode"`
		TaxNumber   string `json:"taxNumber"`
		Address     string `json:"address"`
		HouseNumber string `json:"houseNumber"`
	} `json:"receiverInfo"`
	SenderInfo struct {
		Name        string `json:"name"`
		Company     string `json:"company"`
		CountryId   string `json:"countryId"`
		CountryName string `json:"countryName"`
		Phone       string `json:"phone"`
		State       string `json:"state"`
		City        string `json:"city"`
		Email       string `json:"email"`
		ZipCode     string `json:"zipCode"`
		TaxNumber   string `json:"taxNumber"`
		Address     string `json:"address"`
		HouseNumber string `json:"houseNumber"`
	} `json:"senderInfo"`

	ParcelInfo struct {
		Currency      string `json:"currency"`
		TotalPrice    string `json:"totalPrice"`
		TotalQuantity string `json:"totalQuantity"`
		TotalWeight   string `json:"totalWeight"`
		Height        string `json:"height"`
		Width         string `json:"width"`
		Length        string `json:"length"`
		HasBattery    string `json:"hasBattery"`
		Ioss          string `json:"ioss"`
		ProductList   []struct {
			GoodsNameCh string `json:"goodsNameCh"`
			GoodsNameEn string `json:"goodsNameEn"`
			Quantity    string `json:"quantity"`
			Weight      string `json:"weight"`
			Price       string `json:"price"`
			HsCode      string `json:"hsCode"`
			Url         string `json:"url"`
			Material    string `json:"material"`
		} `json:"productList"`
	} `json:"parcelInfo"`
}
