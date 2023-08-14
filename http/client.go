package http

import (
	"encoding/json"
	"errors"
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
	AccessToken string
	rwLock       sync.RWMutex
}

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 10
)

// IotTokenReq 获取新的token请求参数
type IotTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// IotTokenRefreshReq 刷新token请求参数
type IotTokenRefreshReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// IotTokenRes 获取token请求结果
type IotTokenRes struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Error        string `json:"error"`
}

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

func (s *IotClient) GetNewToken() (res *IotTokenRes, err error) {
	var req = &IotTokenReq{
		GrantType:    "password",
		ClientId:     s.ClientId,
		ClientSecret: s.ClientSecret,
		Username:     s.Username,
		Password:     s.Password,
	}
	response, err := buildHttpClient().
		PostJson(s.AuthUrl+"/api/login/oauth/access_token", req)
	if err != nil {
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	if response.StatusCode != 200 {
		panic(fmt.Sprintf(`iot请求access_token失败：StatusCode = %d`, response.StatusCode))
	}
	bytes, err := response.ReadAll()
	if err != nil {
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	err = json.Unmarshal(bytes, &res)
	if res.Error != "" {
		err = errors.New(fmt.Sprintf("iot请求access_token失败：%s", res.Error))
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	return
}

func (s *IotClient) RefreshToken(scope string, refreshKey string) (res *IotTokenRes, err error) {
	var req = &IotTokenRefreshReq{
		GrantType:    "refresh_token",
		ClientId:     s.ClientId,
		ClientSecret: s.ClientSecret,
		Scope:        scope,
		RefreshToken: refreshKey,
	}
	response, err := buildHttpClient().
		PostJson(s.AuthUrl+"/api/login/oauth/refresh_token", req)
	if err != nil {
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	if response.StatusCode != 200 {
		panic(fmt.Sprintf(`iot请求access_token失败：StatusCode = %d`, response.StatusCode))
	}
	bytes, err := response.ReadAll()
	if err != nil {
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	err = json.Unmarshal(bytes, &res)
	if res.Error != "" {
		err = errors.New(fmt.Sprintf("iot请求access_token失败：%s", res.Error))
		panic(fmt.Sprintf(`iot请求access_token失败：%s`, err.Error()))
	}
	return
}
