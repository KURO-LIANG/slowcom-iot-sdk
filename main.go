package main

import (
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/service"
	"github.com/kuro-liang/slowcom-iot-sdk/config"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

func main() {
	// 调用demo
	client := &http.IotClient{
		BaseUrl:      config.BaseUrl,
		ClientId:     "dee27f91dd6023629ec7",
		ClientSecret: "AE83F1A6F8B14D31EF4B00137756DF5F",
		Username:     "open_fnxgb4",
		Password:     "AE83F1A6F8B14D31EF4B00137756DF5F",
	}
	request := service.GatewayTokenRequest{IotClient: client}

	res, err := request.GetNewToken()
	if err != nil {
		fmt.Println("请求错误：", err)
		return
	}
	fmt.Printf("结果： %+v\n", res)
}
