package entity

// IotAuthSettingReq 第三方回调鉴权配置请求参数
type IotAuthSettingReq struct {
	ClientID  string             `json:"clientId" description:"平台ID"`
	Timestamp string             `json:"timestamp" description:"时间戳"`
	Sign      string             `json:"sign" description:"签名，md5(clientId+timestamp) 加密方式"`
	Data      IotAuthSettingData `json:"data" description:"回调鉴权配置信息"`
}

type IotAuthSettingData struct {
	Domain          string  `json:"domain" description:"鉴权服务器地址"`
	AuthSuccessCode string  `json:"authSuccessCode" description:"鉴权成功返回状态码"`
	AuthFailCode    string  `json:"authFailCode" description:"鉴权失败返回状态码"`
	Timeout         float64 `json:"timeout" description:"超时等待时间(0.1s - 10s)"`
}

// IotAuthSettingRes 第三方回调鉴权配置返回
type IotAuthSettingRes struct {
	ReturnCode    string `json:"return_code" description:"状态码"`
	ReturnMessage string `json:"return_message" description:"错误消息"`
}

// IotThirdAuthFreshReq 第三方刷新用户的鉴权信息请求参数
type IotThirdAuthFreshReq struct {
	ClientID string `json:"clientId" binding:"required" description:"平台ID-即casdoor用户名"`
	Token    string `json:"token" binding:"required" description:"第三方平台用户的登录token"`
	GroupId  string `json:"groupId" binding:"required" description:"组ID"`
}
