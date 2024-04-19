package service

import (
	"fmt"
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

// GetSleepWarningList 获取睡眠带预警配置列表
func (s *HealthRequest) GetSleepWarningList(sn string) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/health/sleep/alarm/setting/list?sn=", sn))
	return
}
