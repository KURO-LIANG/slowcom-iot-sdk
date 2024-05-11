package result

import (
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type SleepAnalysisDetailRes struct {
	http.IotBaseRes
	Data *entity.SleepAnalysisDetail `json:"data" description:"数据"`
}
