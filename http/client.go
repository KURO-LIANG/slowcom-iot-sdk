package http

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	netUrl "net/url"
	"sync"
)

type IotClient struct {
	BaseUrl      string
	AuthUrl      string
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	AccessToken  string
	rwLock       sync.RWMutex
}

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 10
)

// 生成一个http请求客户端
func buildHttpClient() *httpclient.HttpClient {
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		"opt_useragent":      USERAGENT,
		"opt_timeout":        TIMEOUT,
		"Accept-Encoding":    "gzip,deflate,sdch",
		"opt_connecttimeout": CONNECT_TIMEOUT,
		"OPT_DEBUG":          true,
	})
}

// PostJson json请求
func (s *IotClient) PostJson(url string, data interface{}) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.AccessToken).
		PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}

// PutJson json请求
func (s *IotClient) PutJson(url, data interface{}) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.AccessToken).
		PutJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}

// Get get请求
func (s *IotClient) Get(url string) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.AccessToken).
		Get(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}

// Delete 删除请求
func (s *IotClient) Delete(url string) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.AccessToken).Delete(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}
