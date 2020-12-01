package main

import "fmt"

func main() {
	a := 10
	ptr := &a
	fmt.Printf("ptr类型：%T\n", ptr)
	fmt.Printf("ptr指针地址: %p\n", ptr)

	value := *ptr
	fmt.Printf("value值: %d\n", value)
	fmt.Printf("value类型: %T\n", value)
	fmt.Println("--------------------------------")
	c, d := "testDateC", "testDataD"
	swapDate(&c, &d)
	fmt.Println(c, d)
}

func swapDate(c, d *string) {
	swap := *d
	*d = *c
	*c = swap
}
