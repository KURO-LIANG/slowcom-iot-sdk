package service

import (
	"fmt"
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

// DeviceLog 获取设备日志列表 按时间倒叙
func (s *DeviceRequest) DeviceLog(gatewayId string, deviceId string, page int, limit int) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprintf("/device/logs/status/gdid/%s/%s/%d/%d", gatewayId, deviceId, page, limit))
	return
}
