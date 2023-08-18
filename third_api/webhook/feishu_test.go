package webhook

import (
	"fmt"
	"testing"
)

func TestFeiShu_SendMsg(t *testing.T) {
	feiShu := NewFeiShu(&FeiShuConfig{
		URL:  "",
		Sign: "",
	})
	err := feiShu.SendText("测试")
	if err != nil {
		return
	}
	fmt.Println(err)
}
