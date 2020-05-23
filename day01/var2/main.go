// 变量声明
package main

import "fmt"

//var num int 100 //全局作用域
//var num1 = 200 //全局变量
//num3 := 300 //语法错误
func main() {
	/*
		注意点：
		1.变量必须先定义再使用
		2.go语言是静态语言，要求定义变量的类型和赋值是一致的
		3.变量名不能冲突(同一作用域)
		4.简短定义，左边的定量名至少一个是新的
		5.简短定义不能定义全局作用域
		6.变量的零值(初始值)
		7.定义的变量必须使用，不然编译报错
			整形: 0
			浮点型: 0
			字符串: “”
			布尔型: false
	*/
	var n int
	n = 100
	// 打印变量指针地址 %p, &
	fmt.Printf("n的值是：%d,n的内存的地址是：%p\n", n, &n)
	n = 200
	// 打印变量指针地址 %p, &
	fmt.Printf("n的值是：%d,n的内存的地址是：%p\n", n, &n)
	//var s string
	//s = 100  // 变量赋值类型不一致
	//fmt.Println(s) //
	var s string
	s = "李知恩"
	fmt.Println(s)
	n, s = 300, "金泰妍"
	fmt.Println(n, s)
	//fmt.Println(num,num1,num3)
}
