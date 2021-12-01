package time


import (
	"errors"
	"fmt"
	"log"
	"time"
)

func TimeFormatISO8601(timeStr string) time.Time {
	if timeStr != "" || len(timeStr) > 0 {
		newTime, err := time.Parse(Time_TIMEISO8601, timeStr)
		if err != nil {
			log.Println(fmt.Sprintf("【%s】时间格式错误: %v", "base", err))
		}
		return newTime
	} else {
		timeStr = "0001-01-01T00:00:00.000-0700"
		newTime, err := time.Parse(Time_TIMEISO8601, timeStr)
		if err != nil {
			log.Println(fmt.Sprintf("【%s】时间格式错误: %v", "base", err))
		}
		return newTime
	}

}

func TimeFormatStandard(timeStr string) (time.Time, error) {
	if timeStr != "" || len(timeStr) > 0 {
		newTime, err := time.Parse(Time_TIMEStandard, timeStr)
		if err != nil {
			log.Println(fmt.Sprintf("【%s】时间格式错误: %v", "gitlab.hellobike.cn/Defensor/AppContainerManageService", err))
		}
		return newTime, nil
	} else {
		timeStr = "0001-01-01 00:00:00"
		newTime, err := time.Parse(Time_TIMEStandard, timeStr)
		if err != nil {
			log.Println(fmt.Sprintf("【%s】时间格式错误: %v", "base", err))
		}
		return newTime, errors.New("时间格式不正确")
	}

}
