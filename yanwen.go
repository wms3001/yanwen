package yanwen

import (
	"github.com/golang-module/dongle"
	"strconv"
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

func (yanWen *YanWen) Channel() {

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
