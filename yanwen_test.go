package yanwen

import (
	"log"
	"testing"
	"time"
)

var yanwen = YanWen{
	Url:       "https://ejf-fat.yw56.com.cn/",
	UserId:    "100000",
	Format:    "json",
	Method:    "",
	Timestamp: 0,
	Version:   "V1.0",
	ApiToken:  "D6140AA383FD8515B09028C586493DDB",
	Data:      "",
}

func TestYanWen_Sign(t *testing.T) {
	yanwen.Method = "express.channel.getlist"
	yanwen.Timestamp = time.Now().UnixMicro()
	sign := yanwen.Sign()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("Message:")
	log.Println(time.Now().UnixMilli())
	log.Println(sign)
}
