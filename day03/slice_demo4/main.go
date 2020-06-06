package main

import "fmt"

func main() {
	/*
		切片是引用类型数据
		值传递： int float array string
		应用类型数据： slice
	*/

	// 值传递的类型 内存地址会发生改变
	a1 := [...]int{1, 2, 3}
	a2 := a1
	a1[0] = 0
	fmt.Println(a1, a2)
	fmt.Printf("%p\n", &a1)
	fmt.Printf("%p\n", &a2)
	fmt.Println("---------------------")
	//引用类型的 内存地址不会发生改变

	s1 := []int{1, 2, 3}
	s2 := s1

	s1[0] = 0
	fmt.Println(s1, s2)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
}
