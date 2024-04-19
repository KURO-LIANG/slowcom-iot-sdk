package service

import (
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type HealthRequest struct {
	IotClient *http.IotClient
}

// GetSleepReport 获取睡眠报告
func (s *HealthRequest) GetSleepReport(sn string, limit int) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/health/sleep/report/list?sn=", sn, "&limit=", limit))
	return
}

// Subscribe 设置睡眠带预警回调地址
func (s *HealthRequest) Subscribe(req entity.SleepSubscribeReq) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/health/sleep/subscribe", req)
	return
}
