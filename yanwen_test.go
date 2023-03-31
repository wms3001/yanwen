package yanwen

import (
	"github.com/wms3001/yanwen/model"
	"log"
	"testing"
	"time"
)

var yanwen = YanWen{
	Url:       "https://ejf-fat.yw56.com.cn",
	UserId:    "100000",
	Format:    "json",
	Method:    "",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
	Data:      "",
}

//var msg = wlog.Wlog{}

func TestYanWen_Sign(t *testing.T) {
	yanwen.Method = "express.channel.getlist"
	yanwen.Timestamp = time.Now().UnixMicro()
	sign := yanwen.Sign()
	log.Println(sign)
}

func TestYanWen_Channel(t *testing.T) {
	channlResp := yanwen.Channel()
	log.Println(channlResp)
}

func TestYanWen_Country(t *testing.T) {
	countryResp := yanwen.Country()
	log.Println(countryResp)
}

func TestYanWen_Warehouse(t *testing.T) {
	var channlId = model.YanWenWarehouseRequest{}
	channlId.ChannelId = ""
	yanwen.Data = channlId
	warehouseResp := yanwen.Warehouse()
	log.Println(warehouseResp)
	channlId.ChannelId = "481"
	yanwen.Data = channlId
	countryResp1 := yanwen.Warehouse()
	log.Println(countryResp1)
}

func TestYanWen_Order(t *testing.T) {
	var order = model.YanWenOrder{}
	order.OrderNumber = "Test1000000003"
	order.OrderSource = "erp"
	order.UserId = yanwen.UserId
	order.ChannelId = "155"
	order.DateOfReceipt = "2023-03-31"
	order.Remark = ""
	var receiver = model.YanWenAddress{}
	receiver.City = "Los Angeles"
	receiver.Country = "US"
	receiver.Address = "2497  Locust Court"
	receiver.Company = ""
	receiver.Email = "hcixldypft@iubridge.com"
	receiver.HouseNumber = ""
	receiver.Name = "Seratovic"
	receiver.Phone = "424-324-3365"
	receiver.State = "California"
	receiver.TaxNumber = ""
	receiver.ZipCode = "90014"
	order.ReceiverInfo = receiver
	var sender = model.YanWenAddress{}
	sender.City = "ShenZhen"
	sender.Country = "CN"
	sender.Address = "BanTian"
	sender.Company = ""
	sender.Email = ""
	sender.HouseNumber = ""
	sender.Name = "Li Shi"
	sender.Phone = "13324657890"
	sender.State = "GuangDond"
	sender.TaxNumber = ""
	sender.ZipCode = "518110"
	order.SenderInfo = sender
	var parcelInfo = model.YanWenParcel{}
	parcelInfo.Currency = "USD"
	parcelInfo.Height = 5
	parcelInfo.Width = 5
	parcelInfo.Length = 5
	parcelInfo.Ioss = ""
	parcelInfo.HasBattery = 0
	parcelInfo.TotalPrice = 10
	parcelInfo.TotalQuantity = 1
	parcelInfo.TotalWeight = 200
	var productList = []model.YanWenProduct{}
	var product = model.YanWenProduct{}
	product.Url = ""
	product.GoodsNameCh = "上衣"
	product.GoodsNameEn = "shangyi"
	product.Hscode = "1234565437"
	product.Price = 10
	product.Material = ""
	product.Quantity = 1
	product.Weight = 200
	productList = append(productList, product)
	parcelInfo.ProductList = productList
	order.ParcelInfo = parcelInfo
	yanwen.Data = order
	orderResp := yanwen.Order()
	log.Println(orderResp)
}

func TestYanWen_Label(t *testing.T) {
	var labelRequest = model.YanWenLabelRequest{}
	labelRequest.WaybillNumber = "LR039896495CN "
	labelRequest.PrintRemark = 0
	yanwen.Data = labelRequest
	labelResp := yanwen.Label()
	log.Println(labelResp)
}

func TestYanWen_KrPccc(t *testing.T) {
	var krPcccRequest = model.YanWenKrPcccRequest{}
	var receiverInfo = model.YanWenKrPcccReceiverInfo{}
	receiverInfo.Name = "김하성"
	receiverInfo.Phone = "01075654552"
	receiverInfo.TaxNumber = "P972150029777"
	krPcccRequest.ReceiverInfo = receiverInfo
	yanwen.Data = krPcccRequest
	krPcccResp := yanwen.KrPccc()
	log.Println(krPcccResp)
}
