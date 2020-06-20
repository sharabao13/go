package main

import "fmt"

func main() {
	/*
		高阶函数
			根据go语言数据类型特点，可将一个函数作为另一个函数的参数
		fun1(),fun2()
			将fun1 作为fun2这个函数的参数
			fun2 叫高阶函数，接收一个函数作为参数，
			fun 回调函数，作为另一个函数的参数
	*/

	//设计一个函数，用于求加减乘除运算
	fmt.Printf("%T\n", add)  //func(int, int) int
	fmt.Printf("%T\n", oper) //func(int, int, func(int, int) int) int
	fmt.Println(add(1, 3))
	//加法函数调用
	fmt.Println(oper(1, 3, add))

	//减法函数调用
	fmt.Println(oper(10, 7, sub))

	//乘法函数调用
	fmt.Println(oper(20, 10, mul))

	//除法函数调用
	fmt.Println(oper(6, 3, div))

	//乘法函数匿名函数调用
	res2 := oper(2, 5, func(a int, b int) int {
		if a == 0 || b == 0 {
			return 0
		}
		return a * b
	})
	fmt.Println(res2)
}

//加法运算函数
func add(a, b int) int {
	return a + b
}

//减法运算函数
func sub(a, b int) int {
	return a - b
}

//乘法运算函数
func mul(a, b int) int {
	return a * b
}

//除法运算
func div(a, b int) int {
	return a / b
}

func oper(a, b int, fun1 func(int, int) int) int {
	fmt.Println(a, b, fun1) //打印三个参数
	res1 := fun1(a, b)
	return res1
}
