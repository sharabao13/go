package main

import "fmt"

func main() {

	var (
		a int    = 10
		b string = "Tae"
	)
	fmt.Printf("%p %p\n", &a, &b)

	var c string = "Taetae@028 1023 1056"

	// 对字符串取地址
	ptr := &c
	// 打印ptr类型
	fmt.Printf("ptr类型: %T\n", ptr)
	// 打印ptr指针地址
	fmt.Printf("ptr指针地址: %p\n", ptr)

	//对指针地址取值
	value := *ptr
	fmt.Printf("value指针地址的值: %s\n", value)

	//指针地址值的类型
	fmt.Printf("value类型: %T\n", value)
}
