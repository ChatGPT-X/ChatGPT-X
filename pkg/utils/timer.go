package utils

import "time"

// GetCurrentDate 获取当前的时间字符串 格式 "yyyy-mm-dd HH:ii:ss"
func GetCurrentDate() time.Time {
	shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
	currentDateStr := time.Now().Format("2006-01-02 15:04:05")
	currentDate, _ := time.ParseInLocation("2006-01-02 15:04:05", currentDateStr, shanghaiZone)
	return currentDate
}
