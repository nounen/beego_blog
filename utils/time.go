package utils

import "time"

// GetNow 获取 current local time 的地址
func GetNow() *time.Time {
	t := time.Now()
	return &t
}
