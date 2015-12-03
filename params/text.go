package params

////////////////////////////////////////////////////////////////////////////////
// TextCheckParam 百川文档过虑
// http://open.taobao.com/doc2/apiDetail.htm?spm=a219a.7386797.0.0.1u6W4i&apiId=25453
type TextCheckParam struct {
	Content string	// 必须		用户ugc内容
	AppId   string	// 必须		isv的APP对应的key
}

func (this TextCheckParam) APIName() string {
	return "taobao.user.textcheck.get"
}

func (this TextCheckParam) Params() map[string]string {
	var m = make(map[string]string)
	m["content"] = this.Content
	m["app_id"] = this.AppId
	return m
}

func (this TextCheckParam) ExtJSONParamName() string {
	return ""
}

func (this TextCheckParam) ExtJSONParamValue() string {
	return ""
}