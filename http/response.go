package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/kuro-liang/slowcom-iot-sdk/serror"
)

type IotBaseRes struct {
	Code    int    `json:"code" description:"响应码 0-成功"`
	Message string `json:"message" description:"响应消息"`
}
type IotRes struct {
	IotBaseRes
	Data interface{} `json:"data" description:"数据"`
}

// checkResponse 校验请求
func checkResponse(res *httpclient.Response, requestError error) (iotResponse *IotRes, err error) {
	if requestError != nil {
		return nil, serror.New(405, "请求服务异常：请求超时")
	}
	fmt.Printf("iot-sdk 请求结果<== %d - %s\n", res.Response.StatusCode, res.Response.Status)
	if res.Response.StatusCode != 200 {
		return nil, serror.New(res.Response.StatusCode, res.Response.Status)
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &iotResponse)
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	if iotResponse.Code == 0 {
		return
	} else {
		return iotResponse, serror.New(iotResponse.Code, iotResponse.Message)
	}
}
