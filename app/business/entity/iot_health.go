package entity

type SleepRecordResult struct {
	Code         int             `json:"code"`
	ErrorMessage string          `json:"errorMessage"`
	Message      string          `json:"message"`
	Data         []SleepReportV2 `json:"data"`
}

type SleepReport struct {
	ID         int64  `json:"id" description:"ID"`
	GroupId    string `json:"groupId"  description:"组ID"`
	SN         string `json:"sn"  description:"设备ID"`
	UserName   string `json:"userName"  description:"设备绑定的用户名称"`
	GoBedTime  string `json:"goBedTime" description:"就寝时间"`
	OutBedTime string `json:"outBedTime" description:"离床时间"`
	Detail     string `json:"detail" description:"睡眠报告数据"`
	CreateTime string `json:"createTime" description:"报告日期"`
	Remark     string `json:"remark" description:"报告生成的数据分析"`
	WakeUpTime string `json:"wakeUpTime" description:"就寝目标"`
	BedTime    string `json:"bedTime" description:"起床目标"`
}

type SleepReportV2 struct {
	GroupId    string         `json:"groupId"  description:"组ID"`
	SN         string         `json:"sn"  description:"设备ID"`
	CreateTime string         `json:"createTime" description:"报告日期"`
	List       []*SleepReport `json:"list" description:"报告列表，用于版本2，报告按日期分组显示不同时段的报告"`
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
	NoticeMinutes     int    `json:"noticeMinutes" description:"通知时间间隔，单位分钟"`
}

type SleepAnalysisDetail struct {
	GoBedTime         string             `json:"gobedTime" description:"上床时间"`
	OutBedTime        string             `json:"outbedTime" description:"离床时间"`
	DtArr             []string           `json:"dtArr" description:"睡眠时间轴数据集合"`
	SleepHabit        *SleepHabit        `json:"sleepHabit" description:"睡眠习惯"`
	SleepEff          *SleepEff          `json:"sleepEff" description:"睡眠效率"`
	DeepAndLightSleep *DeepAndLightSleep `json:"deepAndLightSleep" description:"深浅睡眠"`
	Stability         *Stability         `json:"stability" description:"安稳度"`
	HeartRate         *HeartRate         `json:"heartRate" description:"心率分析"`
	BreatheRate       *BreatheRate       `json:"breatheRate" description:"呼吸率"`
	AbnormalCondition *AbnormalCondition `json:"abnormalCondition" description:"异常状态"`
}

type AbnormalCondition struct {
	SnoringArr       []int               `json:"snoringArr" description:"打鼾记录"`
	SnoringCountList []*SnoringCountList `json:"snoringCountList" description:"打鼾次数列表"`
	ApneaList        []*ApneaList        `json:"apneaList" description:"疑似呼吸暂停记录"`
	TotalApnea       Apnea               `json:"totalApnea" description:"疑似呼吸暂停汇总"`
	RiskEarlyWarning string              `json:"riskEarlyWarning" description:"风险预警"`
	Knowledge
}

type ApneaList struct {
	Date  string `json:"date" description:"日期"`
	Apnea *Apnea `json:"apnea" description:"呼吸暂停次数"`
}

type Apnea struct {
	Value int    `json:"value" description:"呼吸暂停次数"`
	State int    `json:"state" description:"呼吸暂停评分状态值 1：很差，2-较差，3-正常"`
	Desc  string `json:"desc" description:"呼吸暂停评分状态描述"`
}

type SnoringCountList struct {
	Date         string                        `json:"date" description:"日期"`
	SnoringCount *SmartSleepReportSnoringCount `json:"snoringCount" description:"打鼾次数"`
}
type SmartSleepReportSnoringCount struct {
	Value int    `json:"value" description:"打鼾次数"`
	State int    `json:"state" description:"打鼾评分状态值"`
	Desc  string `json:"desc" description:"打鼾评分状态描述"`
}
type BreatheRate struct {
	AvgHx         SmartSleepReportAvgHx `json:"avgHx" description:"呼吸平均值"`
	HxArr         []int                 `json:"hxArr" description:"呼吸数据集合"`
	MaxMinHistory []*MaxMinHistory      `json:"maxMinHistory" description:"呼吸极值记录"`
	Knowledge
}
type HeartRate struct {
	AvgRh         SmartSleepReportAvgRh `json:"avgRh" description:"心率平均值"`
	RhArr         []int                 `json:"rhArr" description:"心率数据集合"`
	MaxMinHistory []*MaxMinHistory      `json:"maxMinHistory" description:"心率极值记录"`
	Knowledge
}

type MaxMinHistory struct {
	Date     string `json:"date" description:"日期"`
	MaxValue int    `json:"maxValue" description:"最大值"`
	MinValue int    `json:"minValue" description:"最小值"`
	AvgValue int    `json:"avgValue" description:"平均值"`
}

type Stability struct {
	SleepDuration []*SleepDuration            `json:"sleepDuration" description:"睡眠时长列表"`
	OutOfBedTimes []*OutOfBedTimes            `json:"outOfBedTimes" description:"离床次数列表"`
	MoveCountList []*MoveCount                `json:"moveCountList" description:"体动次数列表"`
	RsDuration    RsDuration                  `json:"rsDuration" description:"入睡时长"`
	OutbedCount   SmartSleepReportOutBedCount `json:"outbedCount" description:"离床次数"`
	MoveCount     SmartSleepReportMoveCount   `json:"moveCount" description:"体动次数"`
	Knowledge
}

type SmartSleepReportOutBedCount struct {
	Value int    `json:"value" description:"离床次数"`
	State int    `json:"state" description:"离床评分状态值"`
	Desc  string `json:"desc" description:"离床评分状态描述"`
}

type SmartSleepReportMoveCount struct {
	Value int    `json:"value" description:"体动次数"`
	State int    `json:"state" description:"体动评分状态值"`
	Desc  string `json:"desc" description:"体动评分状态描述"`
}

type RsDuration struct {
	Value int    `json:"value" description:"入睡时长"`
	State int    `json:"state" description:"入睡时长评分状态值 1：较差，2-良好，3-优秀"`
	Desc  string `json:"desc" description:"入睡时长评分状态描述"`
}

type MoveCount struct {
	Date        string `json:"date" description:"日期"`
	BeforeSleep string `json:"beforeSleep" description:"入睡前体动次数"`
	AfterSleep  string `json:"afterSleep" description:"入睡后体动次数"`
}

type OutOfBedTimes struct {
	Date        string `json:"date" description:"日期"`
	BeforeSleep string `json:"beforeSleep" description:"入睡前离床次数"`
	AfterSleep  string `json:"afterSleep" description:"入睡后离床次数"`
}
type SleepDuration struct {
	Date       string `json:"date" description:"日期"`
	RsDuration string `json:"rsDuration" description:"入睡时长"`
}
type DeepAndLightSleep struct {
	AvgRh                    SmartSleepReportAvgRh       `json:"avgRh" description:"心率平均值"`
	SleepArr                 []int                       `json:"sleepArr" description:"睡眠状态数据集合，0-清醒，10-浅睡，20-中睡，30-深睡"`
	DeepAndLightSleepHistory []*DeepAndLightSleepHistory `json:"deepAndLightSleepHistory" description:"深浅睡眠记录"`
	Knowledge
}

type SmartSleepReportAvgRh struct {
	Value int    `json:"value" description:"心率平均值"`
	State int    `json:"state" description:"心率平均值评分值"`
	Desc  string `json:"desc" description:"心率平均值评分描述"`
}
type SmartSleepReportAvgHx struct {
	Value int    `json:"value" description:"呼吸平均值"`
	State int    `json:"state" description:"呼吸平均值评分值"`
	Desc  string `json:"desc" description:"呼吸平均值评分描述"`
}

type DeepAndLightSleepHistory struct {
	Date          string         `json:"date" description:"日期"`
	SoberTime     string         `json:"sober" description:"清醒时长"`
	LightSleep    string         `json:"lightSleep" description:"浅睡时长"`
	MidSleep      string         `json:"midSleep" description:"中睡时长"`
	DeepSleep     string         `json:"deepSleep" description:"深睡时长"`
	DeepSleepRate *DeepSleepRate `json:"deepSleepRate" description:"深睡占比"`
}

type DeepSleepRate struct {
	Value string `json:"value" description:"深睡占比"`
	State int    `json:"state" description:"深睡占比评分状态值 1：较差，2-良好，3-优秀"`
	Desc  string `json:"desc" description:"深睡占比评分状态描述"`
}

type SleepHabit struct {
	List        []*SleepHistory `json:"list" description:"睡眠记录"`
	SleepConfig SleepConfig     `json:"sleepConfig" description:"就寝目标"`
	Knowledge
}
type SleepConfig struct {
	WakeUpTime string `json:"wakeUpTime" binding:"required"  description:"起床目标"`
	BedTime    string `json:"bedTime" binding:"required"  description:"就寝目标"`
}
type Knowledge struct {
	SleepTips       string `json:"sleepTips" description:"睡眠小贴士"`
	HealthKnowledge string `json:"healthKnowledge" description:"健康小知识"`
}

type SleepHistory struct {
	Date       string `json:"date" description:"就寝日期"`
	WakeUpTime string `json:"wakeUpTime" description:"起床时间"`
	BedTime    string `json:"bedTime" description:"就寝时间"`
}

type SleepEff struct {
	List []*SleepEffHistory `json:"list" description:"睡眠效率记录"`
	Knowledge
}

type SleepEffHistory struct {
	Date     string                   `json:"date" description:"日期"`
	SleepEff SmartSleepReportSleepEff `json:"sleepEff" description:"睡眠效率"`
}

type SmartSleepReportSleepEff struct {
	Value string `json:"value" description:"睡眠效率值，百分比，如：0%"`
	State int    `json:"state" description:"睡眠效率评分状态值"`
	Desc  string `json:"desc" description:"睡眠效率评分状态描述"`
}
