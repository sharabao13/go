package main

import "fmt"

func main() {
	/*
		for循环: 某些代码会被多次重复执行
		   语法：
				for 表达式1;表达式2;表达式3; {
						//表达1惯用于变量初始化，执行一次
						//表达2布尔条件循环条件
						//表达3跟着循环体后执行的 用于变量的变化
						循环体
				}
	*/
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("------------------------")

	for i := 0; i < 3; i++ {
		fmt.Println("Hello World")
	}

}
