package params_test

import (
	"fmt"
	"testing"
	"github.com/smartwalle/taobao/params"
	"github.com/smartwalle/taobao"
)

func TestOpenIMAddUserParam(t *testing.T) {
	var p = params.OpenIMAddUserParam{}
	p.UserId = "admin3"
	p.Password = "123456"

	fmt.Println("===== OpenIMAddUserParam =====")
	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
}

func TestOpenIMPushCustomMsgParam(t *testing.T) {
	var p = params.OpenIMPushCustomMsgParam{}
	p.FromUser = "admin2"
	p.ToUsers = []string{"admin"}
	p.Summary = "推送内容"
	p.Data = "消息内容"
	p.SetApsAlert("aa")
	p.SetApsBadge(3)

	fmt.Println("===== OpenIMPushCustomMsgParam =====")
	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
}