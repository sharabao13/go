package main

import "fmt"

//闭包函数

func baseTest(base int) func(int) int {
	return func(n int) int {
		return base + n
	}
}
func main() {
	addBase := baseTest(10)
	v1Base := addBase(8)
	fmt.Println(v1Base)
}
