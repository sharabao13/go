package main

import (
	"fmt"
)

func main() {
	/*
		匿名函数 没有名字的函数
		定义一个匿名函数，直接进行调用，通常只用使用一次，也可使用匿名函数复制给函数变量，就可以使用多次

		匿名函数：
			1.go语言支持函数式编程
			2.将匿名函数作为另一个函数的参数，回调函数
			3.将匿名函数作为另一个函数的返回值，闭包
	*/

	//正常函数调用
	fun1("正常函数")

	//匿名函数创建
	func() {
		fmt.Println("匿名函数")
	}()

	//匿名函数赋值后可多次调用
	fun2 := func() {
		fmt.Println("匿名函数1")
	}
	fun2()

	//带参数的匿名函数
	func(a int) {
		fmt.Println(a)
	}(1)

	//带参数跟返回值的匿名参数
	fun3 := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(fun3)
}

//定义正常的函数

func fun1(s string) {
	fmt.Println(s)
}
