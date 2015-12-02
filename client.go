package taobao

import (
	"fmt"
	"time"
	"sort"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"github.com/smartwalle/going/http"
	"github.com/smartwalle/taobao/params"
)

const (
	TAO_BAO_OPEN_API_URL = "http://gw.api.taobao.com/router/rest"
)

var (
	kTaoBaoAppKey    string
	kTaoBaoAppSecret string
)

func UpdateKey(appKey, appSecret string) {
	kTaoBaoAppKey    = appKey
	kTaoBaoAppSecret = appSecret
}

func Request(param params.IParam) (results map[string]interface{}, err error) {
	results, err = RequestWithKey(kTaoBaoAppKey, kTaoBaoAppSecret, param)
	return results, err
}

func RequestWithKey(appKey, appSecret string, param params.IParam) (results map[string]interface{}, err error) {
	var p = make(map[string]string)
	p["timestamp"]       = time.Now().Format("2006-01-02 15:04:05")
	p["format"]          = "json"
	p["v"]               = "2.0"
	p["sign_method"]     = "md5"
	p["app_key"]         = appKey
	p["method"]          = param.APIName()

	if len(param.ExtJSONParamName()) > 0 {
		p[param.ExtJSONParamName()] = param.ExtJSONParamValue()
	}

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p[key] = value
		}
	}

	var c = http.NewClient()
	c.SetMethod("POST")
	c.SetURLString(TAO_BAO_OPEN_API_URL)

	var keys = make([]string, 0, len(p))
	for key, value := range p {
		c.SetParam(key, value)
		keys = append(keys, key)
	}
	sort.Strings(keys)

	c.SetParam("sign", sign(appSecret, keys, p))

	results, err = c.DoJsonRequest()
	return results, err
}

func sign(appSecret string, keys []string, param map[string]string) (s string) {
	for _, key := range keys {
		s = s + key + param[key]
	}
	s = fmt.Sprintf("%s%s%s", appSecret, s, appSecret)

	var m = md5.New()
	m.Write([]byte(s))
	s = strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
	return s
}