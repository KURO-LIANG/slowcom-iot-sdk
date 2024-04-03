package entity

type ControlContextualReq struct {
	// GroupId 组ID
	// 使用别名时ID此ID必填
	GroupId string `json:"groupId" description:"组ID"`

	// Expiration 过期时间
	// 当为0时为立即执行,不为0时,为离线消息的有效期,将在设备或网关再次上线时执行,
	Expiration int `json:"expiration" description:"过期时间,当为0时为立即执行,不为0时,为离线消息的有效期,将在设备或网关再次上线时执行,"`

	// SceneAlias 情景别名
	SceneAlias string `json:"sceneAlias" description:"情景别名"`

	// SceneId 情景id
	SceneId string `json:"sceneId" description:"情景id"`

	// From 操作来源
	// 格式：平台_用户名
	From string `json:"from" description:"操作来源"`
}
