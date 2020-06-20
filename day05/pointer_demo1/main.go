package main

import "fmt"

func main() {
	/*
		  指针  pointer
				存储另一个变量的内存地址
		  指针定义 语法
				var 变量名 *数据类型
	*/

	a := 10
	fmt.Printf("变量a的数据类型:%T\n ", a)
	fmt.Println("变量a的值是: ", a)
	fmt.Printf("变量a的内存地址是: %p\n", &a)
	fmt.Println("-------------------------")
	// 定义一个指针
	var p1 *int
	fmt.Println(p1) //空指针
	p1 = &a
	fmt.Println("p1指针的数值是：", p1)
	fmt.Printf("p1自身的内存地址是： %p\n", &p1)
	fmt.Println("p1的数值是a的内存地址，该地址存储的数据: ", *p1)

	//操作变量更改数值，并不会改变地址

	a = 200
	fmt.Println(a)
	fmt.Println(&a)

	//通过指针改变数值，

	*p1 = 300
	fmt.Println(*p1)
	fmt.Println(a)
	fmt.Println("--------------------")
	//指针的指针 **
	var p2 **int
	fmt.Println(p2)
	p2 = &p1
	fmt.Printf("%T,%T,%T\n", a, p1, p2)
	fmt.Println("p2的值是: ", p2) //p1的地址
	fmt.Println("p2自身的地址：", &p2)
	fmt.Println("p2中存储的地址 对应的数值 是p1的地址 对应的数值", *p2)
	fmt.Println("p2中存储的地址 对应的数值 再获取对应的数值", **p2)
}
