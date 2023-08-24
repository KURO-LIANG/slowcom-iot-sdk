package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type GatewayTokenRequest struct {
	IotClient *http.IotClient
}

func (s *GatewayTokenRequest) GetNewToken() (res *http.IotRes, err error) {
	var req = &entity.IotTokenReq{
		GrantType:    "password",
		ClientId:     s.IotClient.ClientId,
		ClientSecret: s.IotClient.ClientSecret,
		Username:     s.IotClient.Username,
		Password:     s.IotClient.Password,
	}
	res, err = s.IotClient.PostJson("/api/login/oauth/access_token", req)
	return
}

func (s *GatewayTokenRequest) RefreshToken(scope string, refreshKey string) (res *http.IotRes, err error) {
	var req = &entity.IotTokenRefreshReq{
		GrantType:    "refresh_token",
		ClientId:     s.IotClient.ClientId,
		ClientSecret: s.IotClient.ClientSecret,
		Scope:        scope,
		RefreshToken: refreshKey,
	}
	res, err = s.IotClient.PostJson("/api/login/oauth/refresh_token", req)
	return
}
