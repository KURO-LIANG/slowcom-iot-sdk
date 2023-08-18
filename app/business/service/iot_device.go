package service

import (
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type DeviceRequest struct {
	IotClient *http.IotClient
}

// ModeList 获取设备型号列表
func (s *DeviceRequest) ModeList() (res *http.IotRes, err error) {
	res, err = s.IotClient.Get("/device/model/list")
	return
}
