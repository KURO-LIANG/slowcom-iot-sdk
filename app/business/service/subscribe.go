package service

import (
	"fmt"
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

func (s *GatewaySubscribeRequest) UnSubscribe(groupId string, uuid string) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprintf("/device/subscribeMessages/delete/%s/%s", groupId, uuid))
	return
}
