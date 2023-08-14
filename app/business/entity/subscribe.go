package entity

// SubscribeReq 订阅消息类型请求
type SubscribeReq struct {
	ID          string `json:"id"`                              // 为空时为添加，非空时对应修改
	GroupId     string `json:"groupId" validate:"required"`     // 组ID，当前暂时固定为all
	CallbackUrl string `json:"callbackUrl" validate:"required"` // 回调接口，以POST方式回调
	UUID        string `json:"uuid" validate:"required"`        // 防重复订阅uuid，第三方生成uuid。组id + uuid 会始终唯一，如果重复会被覆盖
	SubContent  []int  `json:"subContent" validate:"required"`  // 订阅类型，订阅类型 -1：全部消息，0-情景执行通知，1-插卡取电(暂未实现)，2-设备下线警告通知,3-SOS报警通知,4-跌倒报警，5-多人聚集，6-燃气报警，7-烟雾报警，8-水浸报警
}
