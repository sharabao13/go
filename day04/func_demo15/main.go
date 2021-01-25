package main

import "fmt"

const (
	//定义每分钟的秒数
	SecondsPerMinute = 60

	//定义每小时的秒数
	SecondsPerHour = 60 * SecondsPerMinute

	//定义每天的秒数
	SecondsPerDay = 24 * SecondsPerHour
)

func resolveTime(seconds int) (minute int, hour int, day int) {
	minute = seconds / SecondsPerMinute
	hour = seconds / SecondsPerHour
	day = seconds / SecondsPerDay
	return
}

func main() {
	fmt.Println(resolveTime(86400))
	// 只获取分钟
	minute, _, _ := resolveTime(86400)
	fmt.Println(minute)
	// 只获取小时
	_, hour, _ := resolveTime(86400)
	fmt.Println(hour)

	_, _, day := resolveTime(86400)
	fmt.Println(day)
}
