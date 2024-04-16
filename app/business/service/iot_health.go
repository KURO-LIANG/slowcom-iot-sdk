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
