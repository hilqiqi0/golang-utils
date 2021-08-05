package net

/*
有关Http协议GET和POST请求的封装
*/
import "github.com/hilqiqi0/golang-utils/config"

type HTTPMethod interface {
	GetHttpConnFromConf(c *config.ConfigEngine, name string)

	SetUrlPara(values ...interface{}) string

	HttpGet(url string) (response string, ok bool)

	HttpPost(url string, data interface{}, contentType string) (content string, err error)
}
