package entity

// IotTokenReq 获取新的token请求参数
type IotTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// IotTokenRefreshReq 刷新token请求参数
type IotTokenRefreshReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// IotTokenRes 获取token请求结果
type IotTokenRes struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Error        string `json:"error"`
}
