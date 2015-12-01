package params

import (
	"encoding/json"
)

type IParam interface{
	// 用于提供访问的 method
	APIName() string

	// 返回请求参数的名字
	ParamName() string

	// 返回参数的 JSON 字符串
	JSONString() string
}


type TaoBaoParam map[string]interface{}

func (this TaoBaoParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

func (this TaoBaoParam) ParamName() string {
	return "send_message_request"
}

func (this TaoBaoParam) JSONString() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this TaoBaoParam) AddParam(key string, value interface{}) {
	this[key] = value
}