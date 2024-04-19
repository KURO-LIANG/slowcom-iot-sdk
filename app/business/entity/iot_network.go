package entity

type AddNetworkReq struct {
	GatewayId       string `json:"gatewayId"`
	GroupId         string `json:"groupId"`
	Alias           string `json:"alias"`
	Model           string `json:"model"`
	ConnectionType  string `json:"connectionType"`
	Type            string `json:"type"`
	Position        string `json:"position"`
	ExpandValue     string `json:"expandValue"`
	SubdivisionType string `json:"subdivisionType"`
	IflyosUserCode  string `json:"iflyosUserCode"`
}
