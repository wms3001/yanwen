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
	yanwen.Data = order
	orderResp := yanwen.Order()
	log.Println(orderResp)
}
