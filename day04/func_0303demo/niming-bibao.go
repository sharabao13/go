package main

import (
	"fmt"
)

func addN(a, b int) int {
	return a + b
}

//callback 格式化函数 将传递的数据按照每行打印还是按照一行按|分割
func print(callback func(...string), args ...string) {
	fmt.Println("print函数输出")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

//闭包结构函数
func increment() func() int {
	//定义一个局部变量
	i := 0
	//定义一个匿名函数给i自增
	fun := func() int {
		i++
		return i
	}
	return fun
}
func main() {
	res1 := increment()
	v1 := res1()
	fmt.Println(v1)
	fmt.Println("--------------------")
	v2 := res1()
	fmt.Println(v2)
	fmt.Println("--------------------")
	fmt.Printf("%T\n", addN)

	res2 := increment()
	v3 := res2()
	fmt.Println(v3)
	fmt.Println("--------------------")
	v4 := res1()
	fmt.Println(v4)
	fmt.Println("--------------------")
	var f func(int, int) int = addN
	fmt.Printf("%T\n", f)

	print(list, "A", "B", "C")
	//匿名函数
	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}
	sayHello("Tae")
	//定义匿名函数并且直接调用

	func(name string) {
		fmt.Println("Hi", name)
	}("Bao")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}, "A", "B", "C")
}
