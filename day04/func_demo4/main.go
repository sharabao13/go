package main

import "fmt"

func main() {
	/*
		可变参数
		一个函数的参数类型确定，个数不确定，可用可变参数

		语法：
			参数名称 ... 参数类型
			对于函数可变参数相当于一个切片
			调用函数时，可出入多个参数
		注意事项：
			1.如果一个函数的参数是可变参数，同时还有其他参数，可变参数要放在末尾
			2.一个函数的可变参数最多只能有一个
	*/

	getAdd(1, 2, 3, 4)

	getAdd(1, 2, 3, 4, -100)
}

//可变参数

func getAdd(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Printf("总和是: %d\n", sum)
}

// 可变参数只能放在末尾
//func fun1(a,b int, strs ... string,c float)  {
//
//}
