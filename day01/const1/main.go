// 常量
package main

import "fmt"

func main() {
	/*
		常量: 常量名一般大写
			1.概念：同变量类似，在程序运行过程中数值不变
			2.语法：
				1.显式类型定义
				2.隐式类型定义
			3.常数
				固定的数值 100，“abc”
	*/
	fmt.Println(100)
	fmt.Println("abc")
	// 显式定义 定义加上类型
	const PATH string = "http://www.google.com"
	// 隐式定义 不带类型
	const PI = 3.141592657
	fmt.Println(PATH, PI)

	// 定义一组常量
	//const C1, C2, C3 = 100, 200, "tae"
	const (
		MALE   = 0
		FEMALE = 1
		UNKOOW = 3
	)
	// 定义一组常量中，某个常量没有赋值，默认和上一行一致
	const (
		A = 100
		B
		C = "ruby"
		D
		E
	)
	fmt.Printf("%T,%d\n", A, A)
	fmt.Printf("%T,%d\n", B, B)
	fmt.Printf("%T,%s\n", C, C)
	fmt.Printf("%T,%s\n", D, D)
	fmt.Printf("%T,%s\n", D, D)
	/*
			iota 特殊常量，可以被编译器修改的常量
		 	每定义一个常量会自动+1
			直到下一个const出现 清零
	*/
	// iota 练习
	const (
		A1 = iota  //0
		B1         //1
		C1 = "mei" //mei //2
		D1         //mei //3
		E1 = iota  //4
		F1 = 100   //100 /5
		G1         //100 /6
	)
	const (
		H1 = iota // 0
	)
	fmt.Println(A1)
	fmt.Println(B1)
	fmt.Println(C1)
	fmt.Println(D1)
	fmt.Println(E1)
	fmt.Println(F1)
	fmt.Println(G1)
	fmt.Println(H1)
}
