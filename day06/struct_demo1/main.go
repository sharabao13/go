package main

import "fmt"

func main() {
	/*
		结构体
			由一系列相同类型或者不同类型的数据的集合
			结构体的成员由一系列的成员变量构成，这些成员别成为字段
	*/

	//初始化结构体
	//方法1
	var p1 Penson
	fmt.Println(p1)
	//结构体访问 .
	p1.name = "mei"
	p1.age = 20
	p1.gender = "male"
	p1.add = "us"
	fmt.Println(p1)

	// 方法2

	p2 := Penson{}
	p2.name = "tae"
	p2.age = 20
	p2.gender = "male"
	p2.add = "zh"
	fmt.Println(p2)

	//方法3

	p3 := Penson{name: "mei1", age: 20, gender: "fmale", add: "zh"}
	fmt.Println(p3)

	// 方法4

	p4 := Penson{"me2", 20, "male", "zh"}
	fmt.Println(p4)

}

// 定义结构体
type Penson struct {
	name   string
	age    int
	gender string
	add    string
}
