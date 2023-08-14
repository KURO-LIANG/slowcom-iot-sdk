package entity

type GatewayGroupAddOrUpdate struct {
	Id    string `json:"id" description:"组ID，无值为添加，有值则是修改"`
	Alias string `json:"alias" description:"组别名"`
	Mode  int    `json:"mode" description:"模式"`
}

type GatewayGroupAddOrUpdateRes struct {
	GroupId string `json:"groupId" description:"组ID"`
}
