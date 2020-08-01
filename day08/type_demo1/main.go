package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		 type 用于类型定义和类型别名
			1.类型定义 type 类型名 type
			2.类型别名 type 类型名=type
	*/
	var i1 myint
	i1 = 100
	i2 := 200
	fmt.Println(i1, i2)

	var s1 mystr
	s1 = "a"
	s2 := "b"
	fmt.Println(s1, s2)
	//s1 = s2 //cannot use s2 (type string) as type mystr in assignment
	fmt.Println(s1)

	res1 := fun1()
	fmt.Println(res1(10, 20))
}

type myint int
type mystr string

//定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}
