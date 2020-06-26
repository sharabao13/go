package main

import (
	"fmt"
)

func main() {
	/*
		匿名结构体和匿名字段
			匿名结构体 没有名字的结构体
				在创建匿名结构体时，同时创建对象

			匿名字段 一个结构的字段没有字段名
	*/

	//匿名函数
	func() {
		fmt.Println("hello world")
	}()

	//结构体初始化
	s1 := Students{"a", 20}
	fmt.Println(s1)

	//匿名结构体
	s2 := struct {
		name string
		age  int
	}{
		name: "b",
		age:  20,
	}
	fmt.Println(s2)

	//结构体匿名字段初始化
	w2 := Worker{"c", 20}
	fmt.Println(w2)
	fmt.Println(w2.string) //结构体匿名字段通过数据类型来访问
}

//普通结构体
type Students struct {
	name string
	age  int
}

//匿名字段

type Worker struct {
	string //字段类型不能重复，否则会冲突
	int
}
