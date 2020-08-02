package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		time 包的使用
	*/
	//1. 获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1) //2020-08-02 14:06:34.7908984 +0800 CST m=+0.002989501

	//2. 获取指定的时间
	t2 := time.Date(2020, 5, 1, 15, 25, 05, 05, time.Local)
	fmt.Println(t2)

	//3.time -- string之间的转换
	//t1.Format 时间模板 日期必须是 2006-1-2 15:04:05
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	//string -- time  parse
	s3 := "1990-05-11 15:08:08"
	t3, err := time.Parse("2006-01-03 15:04:05", s3)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(t3)
	fmt.Println(t1.String())

	// 根据当前时间获取指定的内容
	year, month, date := t1.Date()
	fmt.Println(year, month, date)

	//获取分时秒
	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	//获取星期
	week := t1.Weekday()
	fmt.Println(week)

	// 获取时间戳 1970年1月1日0点0分0秒0纳	秒
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp1 := t4.Unix()
	fmt.Println(timeStamp1)

	t5 := time.Date(2020, 8, 2, 14, 30, 0, 0, time.UTC)
	timeStamp2 := t5.Unix()
	fmt.Println(timeStamp2)
	timeStamp3 := t1.Unix()
	fmt.Println(timeStamp3)

	//时间间隔 add
	t6 := t1.Add(5 * time.Minute)
	fmt.Println(t1)
	fmt.Println(t6)
	t7 := t1.Add(24 * time.Hour)
	fmt.Println(t7)
	fmt.Println(t1)
	fmt.Println("-------------------")

	t8 := t1.AddDate(1, 1, 1)
	fmt.Println(t8)
	fmt.Println(t1)
	fmt.Println("--------------------")
	//时间间隔差值
	t9 := t6.Sub(t1)
	fmt.Println(t9)

	// 睡眠 time.sleep
	time.Sleep(2 * time.Second)
	fmt.Println("main...over...")
	fmt.Println("-----------------")

	// 随机数睡眠[1-10]
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("随机数睡眠...")

}
