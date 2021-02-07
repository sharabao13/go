package main

import "fmt"

func main() {
	var p1 *int
	a := 10
	p1 = &a
	*p1 = 100
	fmt.Println(a, *p1)

	//scan 接收外部输入
	var name string
	fmt.Print("请输入姓名: ")
	fmt.Scan(&name)
	fmt.Printf("您的姓名是: %s\n", name)

	var age int
	fmt.Print("请输入年龄: ")
	fmt.Scan(&age)
	fmt.Printf("您的年龄是: %d\n", age)

	var height float64
	fmt.Print("请输入身高: ")
	fmt.Scan(&height)
	fmt.Printf("您的身高是: %.2f\n", height)

}
