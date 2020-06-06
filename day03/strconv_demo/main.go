package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		 字符串跟基本类型转换
		strconv包
		string convert
	*/

	// 1.bool类型转换
	s1 := "true"
	//将字符串转换bool类型
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T,%t\n", b1, b1)
	}
	fmt.Println("---------------------")
	//将bool类型转换为str类型
	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n", ss1, ss1)
	fmt.Println("---------------------")

	// 2. 整形int
	//strconv.ParseInt(str,base,bitSize) base表示进制，bitsize 表示int类型大写
	s2 := "2"
	i1, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i1, i1)
	fmt.Println("---------------------")
	// 整形转字符串
	ss2 := strconv.FormatInt(i1, 10)
	fmt.Printf("%T,%s\n", ss2, ss2)
	fmt.Println("---------------------")

	//字符串转换整形常用方法atoi
	//整形转字符串itoa

	//字符串转整形
	s5, err := strconv.Atoi("200")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", s5, s5)
	fmt.Println("---------------------")

	//整形转字符串
	i2 := strconv.Itoa(-20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%s", i2, i2)
}
