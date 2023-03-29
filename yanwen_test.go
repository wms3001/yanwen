package yanwen

import (
	"log"
	"testing"
)

func TestYanWen_Sign(t *testing.T) {
	yanwen := YanWen{
		Url:       "",
		UserId:    "",
		Format:    "",
		Method:    "",
		Timestamp: 0,
		Version:   "",
		ApiToken:  "",
		Data:      "",
	}
	sign := yanwen.Sign()
	log.Println(sign)
}
