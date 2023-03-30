package yanwen

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/golang-module/dongle"
	"github.com/wms3001/yanwen/model"
	"log"
	"strconv"
	"time"
)

type YanWen struct {
	Url       string      `json:"url"`
	UserId    string      `json:"userId"`
	Format    string      `json:"format"`
	Method    string      `json:"method"`
	Timestamp int64       `json:"timestamp"`
	Version   string      `json:"version"`
	ApiToken  string      `json:"apiToken"`
	Data      interface{} `json:"data"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("Message:")
}

func (yanWen *YanWen) Channel() model.YanWenChannlResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.channel.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body := []byte{}
	resp := PostReuqest(reqUrl, body)
	var channlResponse model.YanWenChannlResponse
	json.Unmarshal([]byte(resp), &channlResponse)
	return channlResponse
}

func (yanWen *YanWen) Country() model.YanWenCountryResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.country.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body := []byte{}
	resp := PostReuqest(reqUrl, body)
	var countryResponse model.YanWenCountryResponse
	json.Unmarshal([]byte(resp), &countryResponse)
	return countryResponse
}

func (yanWen *YanWen) Warehouse() model.YanWenWarehouseResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "common.warehouse.getlist"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var warehouseResponse model.YanWenWarehouseResponse
	json.Unmarshal([]byte(resp), &warehouseResponse)
	return warehouseResponse
}

func (yanWen *YanWen) Order() model.YanWenOrderResponse {
	yanWen.Timestamp = time.Now().UnixMilli()
	yanWen.Method = "express.order.create"
	reqUrl := yanWen.Url + "/api/order?" +
		"user_id=" + yanWen.UserId +
		"&method=" + yanWen.Method +
		"&format=" + yanWen.Format +
		"&timestamp=" + strconv.FormatInt(yanWen.Timestamp, 10) +
		"&sign=" + yanWen.Sign() +
		"&version=" + yanWen.Version
	body, _ := json.Marshal(yanWen.Data)
	resp := PostReuqest(reqUrl, body)
	var orderResponse model.YanWenOrderResponse
	json.Unmarshal([]byte(resp), &orderResponse)
	return orderResponse
}

func (yanWen *YanWen) Sign() string {
	var signStr string
	if yanWen.Data == "" {
		signStr = yanWen.ApiToken +
			yanWen.UserId +
			yanWen.Format +
			yanWen.Method +
			strconv.FormatInt(yanWen.Timestamp, 10) +
			yanWen.Version +
			yanWen.ApiToken
	} else {
		body, _ := json.Marshal(yanWen.Data)
		signStr = yanWen.ApiToken +
			yanWen.UserId +
			string(body) +
			yanWen.Format +
			yanWen.Method +
			strconv.FormatInt(yanWen.Timestamp, 10) +
			yanWen.Version +
			yanWen.ApiToken
	}
	return dongle.Encrypt.FromString(signStr).ByMd5().ToHexString()
}

func PostReuqest(url string, body []byte) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		//SetResult(). // or SetResult(AuthSuccess{}).
		Post(url)
	if err != nil {
		response := model.YanWenResponse{
			Success: false,
			Code:    "-1",
			Message: err.Error(),
		}
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}
