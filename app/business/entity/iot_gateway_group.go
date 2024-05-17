package entity

type GatewayGroupAddOrUpdate struct {
	Id     string   `json:"id" description:"组ID，无值为添加，有值则是修改"`
	Alias  string   `json:"alias" description:"组别名"`
	Remark string   `json:"remark" description:"备注"`
	Tags   []string `json:"tags" description:"标签 v1版本的model字段可以放这里面来"`
}

type GatewayGroupAddOrUpdateRes struct {
	GroupId string `json:"groupId" description:"组ID"`
}

type GatewayGroupDelete struct {
	Id string `json:"id" description:"组ID"`
}
