package service

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
	"strconv"
	"time"
)

type IotAuthRequest struct {
	IotClient *http.IotClient
}

// AuthSetting 第三方请求回调鉴权配置
func (s *IotAuthRequest) AuthSetting(req entity.IotAuthSettingData) (resData *entity.IotAuthSettingRes, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e9, 10)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(s.IotClient.ClientId+timestamp)))
	var request = entity.IotAuthSettingReq{
		ClientID:  s.IotClient.Username,
		Timestamp: timestamp,
		Sign:      sign,
		Data:      req,
	}
	res, e := s.IotClient.PostJson("/third/conf/view", request)
	d, _ := json.Marshal(res.Data)
	_ = json.Unmarshal(d, &resData)
	err = e
	return
}
