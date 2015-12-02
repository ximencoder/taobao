package taobao

import (
	"testing"
	"fmt"
	"github.com/smartwalle/taobao/params"
)

func send(p params.IParam) {
	fmt.Println(Send("23274732", "****", p))
}

func TestOpenSmsSendMsgParam(t *testing.T) {
	var p = params.TaoBaoOpenSmsSendMsgParam{}
	p.Mobile = "13558858128"
	p.TemplateId = "1025"
	p.AddParam("code", "1234511116")

	fmt.Println("===== 发送短信 =====")
	send(p)
}

func TestCloudPushNoticeIOS(t *testing.T) {
	var p = params.TaoBaoCloudPushNoticeIOSParam{}
	p.Summary = "golanga"
	p.Target = "all"
	p.Environment = "DEV"
	p.TargetValue = "e"
	p.AddParam("badge", 3)

	fmt.Println("===== 向 iOS 设备推送 APNS 消息 =====")
	send(p)
}

func TestCloudPushMessageIOS(t *testing.T) {
	var p = params.TaoBaoCloudPushMessageIOSParam{}
	p.Body = "中文"
	p.Target = "all"
	p.TargetValue = "e"

	fmt.Println("===== 向 iOS 设备推送消息 =====")
	send(p)
}

func TestCloudPush(t *testing.T) {
	var p = params.TaoBaoCloudPushParam{}
	p.Title = "标题"
	p.Body = "hiaa"
	p.DeviceType = 3
	p.Target = "all"
	p.TargetValue = "e"
	p.IOSBadge = "1"
	p.Type = 1
	p.Remind = true
	p.StoreOffline = true
	p.Timeout = 72
	send(p)
}