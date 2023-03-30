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
	Method:    "common.country.getlist",
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
