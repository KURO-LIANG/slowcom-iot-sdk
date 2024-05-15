package service

import (
	"encoding/json"
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/result"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type HealthRequest struct {
	IotClient *http.IotClient
}

// GetSleepReport 获取睡眠报告
func (s *HealthRequest) GetSleepReport(sn string, groupId string, version string, limit int) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/health/sleep/report/list?sn=", sn, "&groupId=", groupId, "&limit=", limit, "&version=", version))
	return
}

// GetSleepWarningList 获取睡眠带预警配置列表
func (s *HealthRequest) GetSleepWarningList(sn string) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/health/sleep/alarm/setting/list?sn=", sn))
	return
}

// AlarmSetting 睡眠带报警设置
func (s *HealthRequest) AlarmSetting(req entity.SleepAlarmSettingReq) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/health/sleep/alarm/setting", req)
	return
}

// AnalysisReportDetail 获取睡眠带报告分析
func (s *HealthRequest) AnalysisReportDetail(id int) (sleepAnalysisDetailRes *result.SleepAnalysisDetailRes, err error) {
	res, err := s.IotClient.Get(fmt.Sprintf("/health/sleep/report/analysis/detail?id=%d", id))
	if err != nil {
		return
	}
	bs, e1 := json.Marshal(res)
	if e1 != nil {
		fmt.Println("解析错误：", e1)
		return
	}
	err = json.Unmarshal(bs, &sleepAnalysisDetailRes)
	if err != nil {
		fmt.Println("解析错误：", err)
		return
	}
	return
}
