package utils

import "time"

// MillTimeStampToTime 毫秒时间戳，转换成秒和纳秒
func MillTimeStampToTime(timestamp int64) time.Time {
	second := timestamp / 1000
	nano := timestamp % 1000 * 1000000
	return time.Unix(second, nano)
}

// SecondTimeStampToTime 秒级时间戳转换成时间格式
func SecondTimeStampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
