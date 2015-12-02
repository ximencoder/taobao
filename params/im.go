package params

import (
	"encoding/json"
)

////////////////////////////////////////////////////////////////////////////////
// OpenIMAddUserParam 添加 IM 用户
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.pMpmGG&apiId=24164&docType=
type OpenIMAddUserParam struct {
	UserId   string   `json:"userid"`               // 必须 im用户名
	Password string   `json:"password"`             // 必须 im密码

	Nick     string   `json:"nick,omitempty"`       // 可选 昵称
	IconURL  string   `json:"icon_url,omitempty"`   // 可选 头像url
	Email    string   `json:"email,omitempty"`      // 可选 email地址
	Mobile   string   `json:"mobile,omitempty"`     // 可选 手机号码
	TaoBaoId string   `json:"taobaoid,omitempty"`   // 可选 淘宝账号
	Remark   string   `json:"remark,omitempty"`     // 可选 remark
	Extra    string   `json:"extra,omitempty"`      // 可选 扩展字段（Json）
	Career   string   `json:"career,omitempty"`     // 可选 职位
	VIP      string   `json:"vip,omitempty"`        // 可选 vip（Json）
	Address  string   `json:"address,omitempty"`    // 可选 地址
	Name     string   `json:"name,omitempty"`       // 可选 名字
	Age      int      `json:"age,omitempty"`        // 可选 年龄
	Gender   string   `json:"gender,omitempty"`     // 可选 性别。M: 男。 F：女
	WeChat   string   `json:"wechat,omitempty"`     // 可选 微信
	QQ       string   `json:"qq,omitempty"`         // 可选 qq
	WeiBo    string   `json:"weibo,omitempty"`      // 可选 微博
}

func (this OpenIMAddUserParam) APIName() string {
	return "taobao.openim.users.add"
}

func (this OpenIMAddUserParam) Params() map[string]string {
	return nil
}

func (this OpenIMAddUserParam) ExtJSONParamName() string {
	return "userinfos"
}

func (this OpenIMAddUserParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMPushMsgParam 发送标准 IM 消息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.f7Uzbf&apiId=25731
type OpenIMPushMsgParam struct {
	FromUser  string    `json:"from_user,omitempty"`    // 可选 消息发送者
	ToUsers   []string  `json:"to_users,omitempty"`     // 可选 消息接受者
	MsgType   int       `json:"msg_type,omitempty"`     // 可选 消息类型。0:文本消息。1:图片消息，只支持jpg、gif。2:语音消息，只支持amr。8:地理位置信息。
	Context   string    `json:"context,omitempty"`      // 可选 发送的消息内容。根据不同消息类型，传不同的值。0(文本消息):填消息内容字符串。1(图片):base64编码的jpg或gif文件。2(语音):base64编码的amr文件。8(地理位置):经纬度，格式如 111,222
	ToAppKey  string    `json:"to_appkey,omitempty"`    // 可选 接收方appkey，默认本app，跨app发送时需要用到
	MediaAttr string    `json:"media_attr,omitempty"`   // 可选 json map，媒体信息属性。根据msgtype变化。0(文本):填空即可。 1(图片):需要图片格式，{"type":"jpg"}或{"type":"gif"}。 2(语音): 需要文件格式和语音长度信息{"type":"amr","playtime":5}
}

func (this OpenIMPushMsgParam) APIName() string {
	return "taobao.openim.immsg.push"
}

func (this OpenIMPushMsgParam) Params() map[string]string {
	return nil
}

func (this OpenIMPushMsgParam) ExtJSONParamName() string {
	return "immsg"
}

func (this OpenIMPushMsgParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMPushCustomMsgParam 发送自定义 IM 消息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.FWTE56&apiId=25185
type OpenIMPushCustomMsgParam struct {
	FromUser  string                    `json:"from_user,omitempty"`    // 可选 消息发送者
	ToUsers   []string                  `json:"to_users,omitempty"`     // 必须 消息接受者
	ToAppKey  string                    `json:"to_appkey,omitempty"`    // 可选 接收方appkey，默认本app，跨app发送时需要用到
	Summary   string                    `json:"summary,omitempty"`      // 必须 客户端最近消息里面显示的消息摘要
	Data      string                    `json:"data,omitempty"`         // 可选 发送的自定义数据，sdk默认无法解析消息，该数据需要客户端自己解析
	Aps       map[string]interface{}    `json:"aps,omitempty"`          // 可选 apns推送时，里面的aps结构体json字符串，aps.alert为必填字段。本字段为可选，若为空，则表示不进行apns推送
	ApsParam  map[string]interface{}    `json:"apns_param,omitempty"`   // 可选 apns推送的附带数据。
}

func (this OpenIMPushCustomMsgParam) APIName() string {
	return "taobao.openim.custmsg.push"
}

func (this OpenIMPushCustomMsgParam) Params() map[string]string {
	return nil
}

func (this OpenIMPushCustomMsgParam) ExtJSONParamName() string {
	return "custmsg"
}

func (this OpenIMPushCustomMsgParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this *OpenIMPushCustomMsgParam) setAps(key string, value interface{}) {
	if this.Aps == nil {
		this.Aps = make(map[string]interface{})
	}
	this.Aps[key] = value
}

func (this *OpenIMPushCustomMsgParam) SetApsAlert(alert string) {
	this.setAps("alert", alert)
}

func (this *OpenIMPushCustomMsgParam) SetApsBadge(badge int) {
	this.setAps("badge", badge)
}

func (this *OpenIMPushCustomMsgParam) SetApsSound(sound string) {
	this.setAps("sound", sound)
}

func (this *OpenIMPushCustomMsgParam) AddApsParam(key string, value interface{}) {
	if this.ApsParam == nil {
		this.ApsParam = make(map[string]interface{})
	}
	this.ApsParam[key] = value
}