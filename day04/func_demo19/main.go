package main

import "fmt"

// 可变参数
// 定义一个或多个参数 一个返回值，返回值为结果相加
func add(a int, arg ...int) int {
	sum := a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}
func strspl(a string, arg ...string) string {
	str := a
	for i := 0; i < len(arg); i++ {
		str += arg[i]
	}
	return str
}

//字符串
func main() {
	fmt.Println(add(8, 9, 10, 11, 12))

	fmt.Println(strspl("guan", "fu", "bao"))
}
