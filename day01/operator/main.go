// 运算符

package main

import "fmt"

func main() {
	/*
		算术运算符：+ - * / % ++ --
		一个表达式不能同时出现 ++ --
	*/
	a := 20
	b := 3

	//加法
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	//减法
	sub := a - b
	fmt.Printf("%d - %d = %d\n", a, b, sub)

	//乘法
	mul := a * b
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	//除法 取商
	div := a / b
	fmt.Printf("%d / %d = %d\n", a, b, div)

	//求余 取模
	mod := a % b
	// %占位符需在前面再加个%
	fmt.Printf("%d %% %d = %d\n", a, b, mod)
}
