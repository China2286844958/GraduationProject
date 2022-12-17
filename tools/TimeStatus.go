package tools

import "time"

/**
@Author 薛智敏
@CreateTime 2022年11月7日23:42:39
*/

// TimeStatus
//
//	@Description: 时间状态
func TimeStatus() (timeStatus string) {

	hour := time.Now().Hour()
	minute := time.Now().Minute()

	//hour := h
	//minute := m
	if hour >= 0 && hour < 6 {
		timeStatus = "凌晨"
	} else if hour >= 6 && hour < 8 {
		timeStatus = "早晨"
	} else if hour >= 8 && hour < 11 {
		timeStatus = "上午"
	} else if hour >= 11 && minute <= 59 && hour < 13 {
		timeStatus = "中午"
	} else if hour >= 13 && hour < 18 {
		timeStatus = "下午"
	} else if hour >= 18 && hour <= 23 {
		timeStatus = "晚上"
	}

	return timeStatus
}
