package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		随机数
		伪随机数 根据一定算法公式推算出来
		math/rand
		没创建种子数每次生成的数字都是固定的 通常跟时间搭配
		设置种子数方法 rand.Seed(int64)
	*/

	//rand.Int() 取随机整数
	num1 := rand.Int()
	fmt.Println("------>", num1)

	//rand.Intn 取范围内的随机数[n) 取值返回0,n-1
	for i := 1; i < 10; i++ {
		num2 := rand.Intn(10)
		fmt.Println("------>", num2)
	}
	rand.Seed(6)
	num3 := rand.Intn(5)
	fmt.Println("------>", num3)

	t1 := time.Now()
	fmt.Println(t1)
	// 时间戳：指定时间 从1970年1月1日过去的时间差值,
	// 秒 Unix()
	// 纳秒 UnixNano()

	timeStmap1 := t1.Unix()
	fmt.Println("------>", timeStmap1)

	timeStmap2 := t1.UnixNano()
	fmt.Println("------>", timeStmap2)

	//设置随机的种子数
	//step1 设置种子数
	rand.Seed(time.Now().UnixNano())
	for j := 1; j < 10; j++ {
		fmt.Println("-->", rand.Intn(10))
	}

	//获取[5,28] 直接的随机数
	num4 := rand.Intn(23) + 5
	fmt.Println(num4)

	//随机种子数，获取[8,66]随机数
	rand.Seed(time.Now().UnixNano())
	for d := 1; d <= 10; d++ {
		fmt.Printf("8,66随机是数：%d\n", rand.Intn(58)+8)
	}
	var b1 int
	fmt.Println("请输入2-100之间的数: ")
	fmt.Scanln(&b1)
	for c1 := 2; c1 < b1; c1++ {
		if b1%c1 == 0 {
			fmt.Printf("%d不是素数", b1)
			break
		} else {
			fmt.Printf("%d是素数", b1)
			break
		}
	}
}
