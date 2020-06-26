package main

import "fmt"

func main() {
	/*
		值类型 int string float bool array struct
		引用类型 map slice fuction pointer
	*/

	p1 := Person{"aaa", 20, "zh"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	p2.name = "bbb"
	fmt.Println(p1)
	fmt.Println(p2)

	// 定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)

	pp1.name = "ccc"
	fmt.Println(*pp1)

	// 使用go内置函数new(),go语言中专门创建某种数据类型的指针的函数
	pp2 := new(Person)
	fmt.Println(pp2)
	fmt.Printf("%p,%T\n", pp2, pp2)
	pp2.name = "ddd"
	pp2.age = 20
	pp2.add = "cn"
	fmt.Println(*pp2)

	//new() 不是nil 空指针
	//指向了新分配的内存空间，里面存储了零值

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
}

type Person struct {
	name string
	age  int
	add  string
}
