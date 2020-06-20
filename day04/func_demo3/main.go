package main

import "fmt"

func main() {
	/*
		多个参数传递
		func funcName(a,b int,string) 同一种数据类型可一起写

		函数调用
		注意点：
			1.函数调用实参跟形参的值必须一一对，顺序 个数 数据类型都要一直
	*/
	getAdd(10, 6)
	fmt.Println("------------------------")

	Calculator(10, "*", 5)
}
func getAdd(a, b int) {
	sum := a + b
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}

//简单计算器

func Calculator(a int, operator string, b int) {
	switch operator {
	case "+":
		sum := a + b
		fmt.Printf("%d+%d=%d\n", a, b, sum)
	case "-":
		sum := a - b
		fmt.Printf("%d-%d=%d\n", a, b, sum)
	case "*":
		sum := a * b
		fmt.Printf("%d*%d=%d\n", a, b, sum)
	case "/":
		sum := a / b
		fmt.Printf("%d/%d=%d\n", a, b, sum)
	case "%":
		sum := a % b
		fmt.Printf("%d%%%d=%d\n", a, b, sum)
	}
}
