package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type GatewaySubscribeRequest struct {
	IotClient *http.IotClient
}

func (s *GatewaySubscribeRequest) Subscribe(req *entity.SubscribeReq) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/device/subscribeMessages/subscribe", req)
	return
}
