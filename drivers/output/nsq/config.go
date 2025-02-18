package nsq

import "github.com/eolinker/eosc"

type Config struct {
	Scopes     []string               `json:"scopes" label:"作用域"`
	Topic      string                 `json:"topic" yaml:"topic" label:"topic"`
	Address    []string               `json:"address" yaml:"address" label:"请求地址"`
	AuthSecret string                 `json:"auth_secret" yaml:"auth_secret" label:"鉴权secret"`
	ClientConf map[string]interface{} `json:"nsq_conf" yaml:"nsq_conf" skip:""`
	Type       string                 `json:"type" yaml:"type" enum:"json,line" label:"输出格式"`
	Formatter  eosc.FormatterConfig   `json:"formatter" yaml:"formatter" label:"格式化配置"`
}
