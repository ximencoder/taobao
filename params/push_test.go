package params_test

//import (
//	"fmt"
//	"testing"
//	"github.com/smartwalle/taobao/params"
//	"github.com/smartwalle/taobao"
//)

//func TestCloudPushNoticeToIOSParam(t *testing.T) {
//	var p = params.CloudPushNoticeToIOSParam{}
//	p.Summary = "golang"
//	p.Target = "all"
//	p.Environment = "DEV"
//	p.TargetValue = "e"
//	p.AddParam("badge", 3)
//
//	fmt.Println("===== TestCloudPushNoticeToIOSParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestCloudPushMessageToIOSParam(t *testing.T) {
//	var p = params.CloudPushMessageToIOSParam{}
//	p.Body = "中文"
//	p.Target = "all"
//	p.TargetValue = "e"
//
//	fmt.Println("===== TestCloudPushMessageToIOSParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}
//
//func TestCloudPushParam(t *testing.T) {
//	var p = params.CloudPushParam{}
//	p.Title = "标题"
//	p.Body = "hiaa"
//	p.DeviceType = 0
//	p.Target = "all"
//	p.TargetValue = "e"
//	p.IOSBadge = "1"
//	p.Type = 0
//	p.Remind = true
//	p.StoreOffline = true
//	p.Timeout = 3
//
//	fmt.Println("===== TestCloudPushParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}