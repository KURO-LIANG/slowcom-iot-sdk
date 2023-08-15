package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type IotAuthRequest struct {
	IotClient *http.IotClient
}

// AuthSetting 第三方请求回调鉴权配置
func (s *IotAuthRequest) AuthSetting(req *entity.IotAuthSettingReq) (resData *entity.IotAuthSettingRes, err error) {
	res, err := s.IotClient.PostJson("/v1/api/third/conf/view", req)
	if err == nil {
		resData = res.Data.(*entity.IotAuthSettingRes)
	}
	return
}
