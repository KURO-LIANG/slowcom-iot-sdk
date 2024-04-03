package entity

// GatewayControlForm 网关控制表单结构体
type GatewayControlForm struct {
	// GatewayId 网关ID,WIFI直连设备ID
	// gatewayAlias与gatewayId二选一
	GatewayId string `json:"gatewayId" description:"网关ID,WIFI直连设备ID，gatewayAlias与gatewayId二选一"`

	// DeviceId 设备ID,WIFI直连设备ID-X
	// deviceAlias与deviceId二选一
	DeviceId string `json:"deviceId" description:"设备ID,WIFI直连设备ID-X，deviceAlias 与 deviceId 二选一"`

	// Value 取值
	// 开关类或未注明分类
	// 通常0为关,1为开
	// 门锁
	// 锁开锁时传开门密码
	// 红外转发器
	// 自定义码
	// 发射,{irCodeId}
	// 别名控制
	// 别名兼容通用控制
	// 此时的别名应该为 红外转发器别名-红外控制设备别名 两者之间使用”-“链接
	Value string `json:"value" description:"value取值"`

	// GatewayAlias 网关别名
	// gatewayAlias与gatewayId二选一
	// gatewayId==NULL 或空字符时生效
	GatewayAlias string `json:"gatewayAlias" description:"网关别名 gatewayAlias与gatewayId二选一"`

	// DeviceAlias 设备别名
	// deviceAlias与deviceId二选一
	DeviceAlias string `json:"deviceAlias" description:"设备id deviceAlias 与 deviceId 二选一"`

	// GroupId （家庭）组ID
	// 网关ID和设备ID存在一个未填时，此ID必填
	GroupId string `json:"groupId" description:"组ID，网关ID和设备ID存在一个未填时，此ID必填"`

	// Expiration 过期时间
	// 当为0时为立即执行,不为0时,为离线消息的有效期,将在设备或网关再次上线时执行
	Expiration int `json:"expiration" description:"过期时间,当为0时为立即执行,不为0时,为离线消息的有效期,将在设备或网关再次上线时执行"`

	// From 操作来源
	// 格式：平台_用户名
	From string `json:"from" description:"操作来源" notes:"格式：平台_用户名"`
}
