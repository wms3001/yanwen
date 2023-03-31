# 燕文物流

------

实现了燕文物流各个接口的请求

------

1. 安装
```go
go get github.com/wms3001/yanwen
```
2. 签名算法
```go
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
yanwen.Method = "express.channel.getlist"
yanwen.Timestamp = time.Now().UnixMicro()
sign := yanwen.Sign()
```
3. 获取渠道
```go
var yanwen = YanWen{
	Url:       "https://ejf-fat.yw56.com.cn",
	UserId:    "100000",
	Format:    "json",
	Method:    "express.channel.getlist",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
	Data:      "",
}
channlResp := yanwen.Channel()
```
4. 获取国家
```go
var yanwen = YanWen{
	Url:       "https://ejf-fat.yw56.com.cn",
	UserId:    "100000",
	Format:    "json",
	Method:    "common.country.getlist",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
	Data:      "",
}
countryResp := yanwen.Country()
```
5. 获取仓库
```go
var yanwen = YanWen{
	Url:       "https://ejf-fat.yw56.com.cn",
	UserId:    "100000",
	Format:    "json",
	Method:    "",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
}

//不传channlId 
var channlId = model.YanWenWarehouseRequest{}
channlId.ChannelId = ""
yanwen.Data = channlId
warehouseResp := yanwen.Warehouse()
//传channlId
channlId.ChannelId = "481"
yanwen.Data = channlId
countryResp1 := yanwen.Warehouse()
```
6. 创建订单
```go
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
```
7. 获取面单
```go
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
var labelRequest = model.YanWenLabelRequest{}
labelRequest.WaybillNumber = "LR039896495CN "
labelRequest.PrintRemark = 0
yanwen.Data = labelRequest
labelResp := yanwen.Label()
```
8. 韩国个人通关码校验
```go
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
var krPcccRequest = model.YanWenKrPcccRequest{}
var receiverInfo = model.YanWenKrPcccReceiverInfo{}
receiverInfo.Name = "김하성"
receiverInfo.Phone = "01075654552"
receiverInfo.TaxNumber = "P972150029777"
krPcccRequest.ReceiverInfo = receiverInfo
yanwen.Data = krPcccRequest
krPcccResp := yanwen.KrPccc()
```