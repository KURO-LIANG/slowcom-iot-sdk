package main

import (
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/service"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

func main() {
	// 调用demo
	//client := &http.IotClient{
	//	BaseUrl:      config.BaseUrl,
	//	ClientId:     "dee27f91dd6023629ec7",
	//	ClientSecret: "AE83F1A6F8B14D31EF4B00137756DF5F",
	//	Username:     "open_fnxgb4",
	//	Password:     "AE83F1A6F8B14D31EF4B00137756DF5F",
	//}
	//request := service.GatewayTokenRequest{IotClient: client}
	//
	//res, err := request.GetNewToken()
	//fmt.Printf("结果： %+v\n", res)

	client := &http.IotClient{
		BaseUrl:      "https://control.iot.slowcom.cn/v1/api",
		ClientId:     "dee27f91dd6023629ec7",
		ClientSecret: "AE83F1A6F8B14D31EF4B00137756DF5F",
		Username:     "open_fnxgb4",
		Password:     "AE83F1A6F8B14D31EF4B00137756DF5F",
		AccessToken:  "eyJhbGciOiJSUzI1NiIsImtpZCI6ImNlcnRfZ2RsenU2IiwidHlwIjoiSldUIn0.eyJvd25lciI6IlNsb3djb20iLCJuYW1lIjoib3Blbl9mbnhnYjQiLCJpc3MiOiJodHRwczovL2F1dGguc2xvd2NvbS5jbiIsInN1YiI6ImJjNDkxMDc1LTI2MDUtNDM1Mi04MTkwLTlkYWZkZjYwMzAyMyIsImF1ZCI6WyJkZWUyN2Y5MWRkNjAyMzYyOWVjNyJdLCJleHAiOjE2OTM0NjkxODQsIm5iZiI6MTY5Mjg2NDM4NCwiaWF0IjoxNjkyODY0Mzg0fQ.kUdviZLiDkjq6Mlm1VNUDUgi8WFfjZOLMRWtuz0Nsgu5xAW7Lw9U1XOCaRMq5kI_c_xvmqzV_vC4Lux75lPei9SQByVfGl6ygdbR7k0Gp7Klbl0BPbapZVBgF-m7tfRiIaw5sWAQ2mACQeifEjnz27769Vb5xNCmVWsd3EaQcYMJMLo3WU0g56nOo_fFQ0L7HiOzqdqMSaFK9V0VUKkt0uiEjUlEwrt2bmWimyLu2OGmPxh3Y8lV7jQ7GIalfAvsRHhklEGN9ExGPe74l2g1y5AY2Tg5OvEN7pkwXKhWWyAAFXwUe60GBg_v_KdTwxBo5YbOTfFA2ew5jBZjTFgU7KDHlIc4wIal4hzAVQb09CaWuGOKtdY9CU1zYkQqqjtiP0dW7FCJxr0NfgcCesvlGt491M2iK-FaIevPUgTop4DcbdvpjnVvpX-g5IUgil0Bl-d8fonyAiGBBR1LVSSoQNJCsqdtd-89SaSCFyrPa0ebfco19J4TZHdRDN7EvB-pkaa78qawXU7pnnGrnlF3ZDUvctFwebOKgYApUc71e9MBh0r41mwfJwlq-A5USf2Hi1OGH3_BnUT27il8Etns-1HL-3sokjQZB2in_qgP4q8W3zSWnqS5Sx3CYaWrSIBdi5Ps27FdYSoHNlBELpZEfqNzbAtpQmS660H0KsSgOJk",
	}
	request1 := service.GatewayGroupRequest{IotClient: client}
	//
	res1, err := request1.Save(&entity.GatewayGroupAddOrUpdate{
		Id:    "64e71104d1e20900011f34ea",
		Alias: "IOT SDK测试接口11111",
		Mode:  0,
	})
	if err != nil {
		fmt.Println("请求错误：", err)
		return
	}
	fmt.Printf("结果： %+v\n", res1)
}
