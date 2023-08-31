package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/serror"
	netUrl "net/url"
	"sync"
)

type IotClient struct {
	BaseUrl      string
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
	CONNECT_TIMEOUT = 30
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
	res, err := buildHttpClient().WithHeader("Authorization", s.AccessToken).
		PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}

// PutJson json请求
func (s *IotClient) PutJson(url, data interface{}) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", s.AccessToken).
		PutJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}

// Get get请求
func (s *IotClient) Get(url string) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", s.AccessToken).
		Get(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}

// Delete 删除请求
func (s *IotClient) Delete(url string) (response *IotRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", s.AccessToken).Delete(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}

// GetToken 获取token
func (s *IotClient) GetToken(url string, data interface{}) (response *entity.IotTokenRes, err error) {
	res, err := buildHttpClient().PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return nil, serror.New(405, "请求服务异常：请求超时")
	}
	if res.Response.StatusCode != 200 {
		return nil, serror.New(res.Response.StatusCode, fmt.Sprint(res.Response.StatusCode, " 请求服务异常"))
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	if response.Error == "" {
		return
	} else {
		return response, serror.New(500, response.Error)
	}
}
