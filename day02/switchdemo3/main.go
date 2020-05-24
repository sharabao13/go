package main

import "fmt"

func main() {
	// switch 每月天数打印练习
	// 1,3,5,7,8,10,12 31天
	// 4,6,9,11        30
	// 2闰年 29 天 闰年 能被4整除不能被100整除 或能被400整除

	var (
		year  int
		month int
		day   int
	)
	fmt.Println("请输入一个年份：")
	fmt.Scanln(&year)

	fmt.Println("请输入一个月份: ")
	fmt.Scanln(&month)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("无效月份")
	}
	fmt.Printf("%d 年 %d 月的天数是: %d", year, month, day)
}
