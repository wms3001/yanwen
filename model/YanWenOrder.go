package model

type YanWenOrder struct {
	ChannelId     string        `json:"channelId"`
	OrderSource   string        `json:"orderSource"`
	UserId        string        `json:"userId"`
	OrderNumber   string        `json:"orderNumber"`
	DateOfReceipt string        `json:"dateOfReceipt"`
	Remark        string        `json:"remark"`
	ReceiverInfo  YanWenAddress `json:"receiverInfo"`
	SenderInfo    YanWenAddress `json:"senderInfo"`
	ParcelInfo    YanWenParcel  `json:"parcelInfo"`
}
