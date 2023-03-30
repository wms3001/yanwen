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
	Url       string `json:"url"`
	UserId    string `json:"userId"`
	Format    string `json:"format"`
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
	ApiToken  string `json:"apiToken"`
	Data      string `json:"data"`
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
		signStr = yanWen.ApiToken +
			yanWen.UserId +
			yanWen.Data +
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
