package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type GatewayGroupRequest struct {
	IotClient *http.IotClient
}

// Save 添加或修改组
func (s *GatewayGroupRequest) Save(req *entity.GatewayGroupAddOrUpdate) (resData *http.IotRes, err error) {
	resData, err = s.IotClient.PostJson("/device/group/split/save", req)
	return
}
