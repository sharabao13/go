// 比较运算符
package main

import "fmt"

func main() {
	/*
		比较运算符：> < >= <= == !=
		返回值 布尔类型
	*/
	a := 6
	b := 8
	c := 6
	res1 := a > b  // f
	res2 := b > c  // t
	res3 := a == b // f
	res4 := a == c // t
	res5 := a != b // t
	res6 := a != c // f
	fmt.Printf("%T,%t\n", res1, res1)
	fmt.Printf("%T,%t\n", res2, res2)
	fmt.Printf("%T,%t\n", res3, res3)
	fmt.Printf("%T,%t\n", res4, res4)
	fmt.Printf("%T,%t\n", res5, res5)
	fmt.Printf("%T,%t\n", res6, res6)
}
