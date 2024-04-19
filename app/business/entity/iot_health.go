package entity

type SleepReport struct {
	ID         int    `json:"id"`
	SN         string `json:"sn"`
	UserName   string `json:"userName"`
	CreateTime string `json:"createTime"`
	Detail     Detail `json:"detail"`
}

type Detail struct {
	Score         Count    `json:"score"`
	TotalDuration int      `json:"total_duration"`
	InBedDuration int      `json:"inbed_duration"`
	GoBedTime     string   `json:"gobed_time"`
	OutBedTime    string   `json:"outbed_time"`
	OutBedCount   Count    `json:"outbed_count"`
	MoveCount     Count    `json:"move_count"`
	SnoringCount  Count    `json:"snoring_count"`
	SleepEff      Count    `json:"sleep_eff"`
	DtArr         []string `json:"dt_arr"`
	RhArr         []int    `json:"rh_arr"`
	HxArr         []int    `json:"hx_arr"`
	SnoringArr    []int    `json:"snoring_arr"`
	OutBedArr     []int    `json:"outbed_arr"`
	MoveArr       []int    `json:"move_arr"`
	HxStopArr     []int    `json:"hxstop_arr"`
	SleepArr      []int    `json:"sleep_arr"`
}

type Count struct {
	Value int    `json:"value"`
	State int    `json:"state"`
	Desc  string `json:"desc"`
}

type SleepSubscribeReq struct {
	SN          string `json:"sn"`
	CallbackUrl string `json:"callbackUrl"`
	UUID        string `json:"uuid"`
}
