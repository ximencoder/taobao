package params

import (
	"encoding/json"
)

type TaoBaoOpenSmsSendMsgParam struct {
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

func (this TaoBaoOpenSmsSendMsgParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

func (this TaoBaoOpenSmsSendMsgParam) Params() map[string]string {
	return nil
}

func (this TaoBaoOpenSmsSendMsgParam) ExtJSONParamName() string {
	return "send_message_request"
}

func (this TaoBaoOpenSmsSendMsgParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this *TaoBaoOpenSmsSendMsgParam) AddParam(key string, value interface{}) {
	if this.Context == nil {
		this.Context = make(map[string]interface{})
	}
	this.Context[key] = value
}