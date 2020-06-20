package main

import "fmt"

func main() {
	/*
		return 语句
			1.一个函数定义了返回值，必须使用return语句返回给调用处，retrun后的数据需和定义的顺序，个数，类型一致
			2.可用使用_，舍去多余的返回值，
			3.如果一个函数定义了有返回值，函数中的分支，循环，需保证无论哪个分支，循环，都有要return语句被执行
			4，如果一个函数没有定义返回值，可用在函数分支中定义return，专门用于函数结束执行
	*/
	// 调用fun1函数
	a, b, c := fun1()
	fmt.Println(a, b, c)

	//可以使用_舍去不用的返回值
	_, _, d := fun1()
	fmt.Println(d)

	//分支内的return调用
	age1 := fun2(-10)
	fmt.Println(age1)

	//
	fun3()
}

// 创建一个多返回值函数

func fun1() (int, float64, string) {
	return 1, 3.1, "Hello"
}

// 分支return
func fun2(age int) int {
	if age > 0 {
		return age
	} else {
		return 0
	}
}

//结束分支retrun
func fun3() {
	for i := 1; i <= 9; i++ {
		if i == 5 {
			//break 跳出该循环继续执行下面语句
			return //结束函数执行，
		}
		fmt.Println(i)
	}
	fmt.Println("fun3...main...over...")
}
