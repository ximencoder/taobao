package params_test

//import (
//	"fmt"
//	"testing"
//	"github.com/smartwalle/taobao/params"
//	"github.com/smartwalle/taobao"
//)
//
//func TestOpenIMAddUsersParam(t *testing.T) {
//	var p = params.OpenIMAddUsersParam{}
//
//	var u1 = params.OpenIMUser{}
//	u1.UserId = "admin4"
//	u1.Password = "123456"
//
//	var u2 = params.OpenIMUser{}
//	u2.UserId = "admin5"
//	u2.Password = "123456"
//
//	p.AddOpenIMUser(&u1)
//	p.AddOpenIMUser(&u2)
//
//	fmt.Println("===== OpenIMAddUserParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestOpenIMUpdateUsersParam(t *testing.T) {
//	var p = params.OpenIMUpdateUsersParam{}
//
//	var u1 = params.OpenIMUser{}
//	u1.UserId = "admin4"
//	u1.Password = "123456"
//
//	var u2 = params.OpenIMUser{}
//	u2.UserId = "admin5"
//	u2.Password = "123456"
//
//	p.AddOpenIMUser(&u1)
//	p.AddOpenIMUser(&u2)
//
//	fmt.Println("===== OpenIMUpdateUsersParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestOpenIMGetUsersParam(t *testing.T) {
//	var p = params.OpenIMGetUsersParam{}
//	p.UserIds = []string{"admin", "admin2"}
//	fmt.Println("===== OpenIMDeleteUserParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestOpenIMDeleteUsersParam(t *testing.T) {
//	var p = params.OpenIMDeleteUsersParam{}
//	p.UserIds = []string{"admin", "admin2"}
//	fmt.Println("===== OpenIMDeleteUserParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestOpenIMPushCustomMsgParam(t *testing.T) {
//	var p = params.OpenIMPushCustomMsgParam{}
//	p.FromUser = "admin2"
//	p.ToUsers = []string{"admin"}
//	p.Summary = "推送内容"
//	p.Data = "消息内容"
//	p.SetApsAlert("aa")
//	p.SetApsBadge(3)
//
//	fmt.Println("===== OpenIMPushCustomMsgParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}