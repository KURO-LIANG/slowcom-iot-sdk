package entity

type SleepRecordResult struct {
	Code         int           `json:"code"`
	ErrorMessage string        `json:"errorMessage"`
	Message      string        `json:"message"`
	Data         []SleepReport `json:"data"`
}

type SleepReport struct {
	ID         int    `json:"id"`
	SN         string `json:"sn"`
	UserName   string `json:"userName"`
	CreateTime string `json:"createTime"`
	Detail     string `json:"detail"`
}

type Detail struct {
	Score         Count       `json:"score"`
	TotalDuration int         `json:"total_duration"`
	InBedDuration int         `json:"inbed_duration"`
	GoBedTime     string      `json:"gobed_time"`
	OutBedTime    string      `json:"outbed_time"`
	OutBedCount   Count       `json:"outbed_count"`
	MoveCount     Count       `json:"move_count"`
	SnoringCount  Count       `json:"snoring_count"`
	SleepEff      interface{} `json:"sleep_eff"`
	DtArr         []string    `json:"dt_arr"`
	RhArr         []int       `json:"rh_arr"`
	HxArr         []int       `json:"hx_arr"`
	SnoringArr    []int       `json:"snoring_arr"`
	OutBedArr     []int       `json:"outbed_arr"`
	MoveArr       []int       `json:"move_arr"`
	HxStopArr     []int       `json:"hxstop_arr"`
	SleepArr      []int       `json:"sleep_arr"`
}

type Count struct {
	Value int    `json:"value"`
	State int    `json:"state"`
	Desc  string `json:"desc"`
}

type SleepAlarmSettingReq struct {
	Sn   string               `json:"sn" binding:"required" description:"设备SN"`
	List []*SleepAlarmSetting `json:"list" binding:"required" description:"报警设置列表"`
}
type SleepAlarmSetting struct {
	AlarmType         int    `json:"alarmType" description:"设置类型 0：心率异常提醒 1：呼吸率异常提醒 2：离床预警提醒 3：紧急求助开关"` // 设置类型 0：心率异常提醒 1：呼吸率异常提醒 2：离床预警提醒 3：紧急求助开关
	MinSpeed          uint32 `json:"minSpeed" description:"最小值"`                                        // 最小值
	MaxSpeed          uint32 `json:"maxSpeed" description:"最大值"`                                        // 最大值
	OutOfBedStartTime string `json:"outOfBedStartTime" description:"离床预警区间-开始时间"`                       // 离床预警区间-开始时间
	OutOfBedEndTime   string `json:"outOfBedEndTime" description:"离床预警区间-结束时间"`                         // 离床预警区间-结束时间
	OutOfBedInterval  int    `json:"outOfBedInterval" description:"离床未归提醒分钟数"`
}
