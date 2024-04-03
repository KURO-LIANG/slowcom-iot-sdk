package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type SceneRequest struct {
	IotClient *http.IotClient
}

// ControlContextual 控制情景
func (s *SceneRequest) ControlContextual(req entity.ControlContextualReq) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/scene/control", req)
	return
}
