package service

import (
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type SceneRequest struct {
	IotClient *http.IotClient
}

// ControlContextual 控制情景
func (s *SceneRequest) ControlContextual(req entity.ControlContextualReq) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/device/scene/control", req)
	return
}

// SceneList 查询组的情景列表
func (s *SceneRequest) SceneList(groupId string, isPrivate uint8) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprintf("/device/scene/detail/%s/%d", groupId, isPrivate))
	return
}
