package main

import "fmt"

func main() {
	/*
	 if 练习
	*/
	var (
		a   int
		b   int
		res int
		ope string
	)
	fmt.Println("----------------------------")

	fmt.Println("请输入第一个整数： ")
	fmt.Scanln(&a)

	fmt.Println("请输入第二个整数： ")
	fmt.Scanln(&b)

	fmt.Println("请输入一个操作符(+-*/%)： ")
	fmt.Scanln(&ope)

	if ope == "+" {
		res = a + b
		fmt.Printf("%d + %d = %d\n", a, b, res)
	} else if ope == "-" {
		res = a - b
		fmt.Printf("%d - %d = %d\n", a, b, res)
	} else if ope == "*" {
		res = a * b
		fmt.Printf("%d * %d = %d\n", a, b, res)
	} else if ope == "/" {
		res = a / b
		fmt.Printf("%d / %d = %d\n", a, b, res)
	} else if ope == "%" {
		res = a % b
		fmt.Printf("%d %% %d = %d\n", a, b, res)
	} else {
		fmt.Println("请输入正确的整数或操作符")
	}
	fmt.Println("----------------------------")
}
