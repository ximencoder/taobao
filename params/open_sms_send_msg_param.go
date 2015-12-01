package params

import (
	"fmt"
	"encoding/json"
)

type OpenSmsSendMsgParam struct {
	// 必须参数
	TemplateId            string                    `json:"template_id,omitempty"`
	Context               map[string]interface{}    `json:"context"`
	Mobile                string                    `json:"mobile"`

	// 可选参数
	SignatureId           string                    `json:"signature_id,omitempty"`
	ExternalId            string                    `json:"external_id,omitempty"`
	DeviceLimit           int                       `json:"external_id,omitempty"`
	SessionLimit          int                       `json:"session_limit,omitempty"`
	DeviceLimitInTime     int                       `json:"device_limit_in_time,omitempty"`
	Mobile_limit          int                       `json:"mobile_limit,omitempty"`
	SessionLimitInTime    int                       `json:"session_limit_in_time,omitempty"`
	MobileLimitInTime     int                       `json:"mobile_limit_in_time,omitempty"`
	SessionId             string                    `json:"session_id,omitempty"`
	Domain                string                    `json:"domain,omitempty"`
	DeviceId              string                    `json:"device_id,omitempty"`
}

func (this OpenSmsSendMsgParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

func (this OpenSmsSendMsgParam) ParamName() string {
	return "send_message_request"
}

func (this OpenSmsSendMsgParam) JSONString() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func (this *OpenSmsSendMsgParam) AddParam(key string, value interface{}) {
	if this.Context == nil {
		this.Context = make(map[string]interface{})
	}
	this.Context[key] = value
}