package main

import (
	"fmt"
)

func main() {
	/*
			函数
			执行特定任务的代码块
			Go 语言最少有个 main() 函数。

			你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

			函数声明告诉了编译器函数的名称，返回类型，和参数。

			Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。
			如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

			语法：
				func function_name( [parameter list] ) [return_types] {
		   函数体
		}

			func：函数由 func 开始声明
			function_name：函数名称，函数名和参数列表一起构成了函数签名。
			parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
			return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
			函数体：函数定义的代码集合。
	*/

	// 求1-100的整数和
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Printf("1-100的和是： %d\n", sum)

	//调用1-100的和的函数
	//调用语法 函数名+（）
	getSUm()
	fmt.Println("---------------------")

	//调用99乘法表
	multiNine()
	fmt.Println("---------------------")
	sayHi("Tae")
	sayHi("Bao")
	sums := addN(1, 2, 5, 2)
	fmt.Println(sums)
	sum1 := calc("add", 1, 2, 5, 6)
	fmt.Println(sum1)
	nums := []int{1, 2, 5, 6, 10}
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
	nums1 := []int{1, 3, 5, 7, 9, 11, 13, 22, 33, 55}
	nums1 = append(nums1[:5], nums1[6:]...)
	fmt.Println(nums1)
}

//定义一个1-100的和的函数
//声明一个 getSum的函数名
func getSUm() {
	//函数体，代码块
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1-100的和是： %d\n", sum)
}

//99乘除法表函数
func multiNine() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func sayHi(name string) {
	fmt.Println("你好", name)
}

// 可变参数的类型为 切片类型
func addN(a, b int, args ...int) int {
	total := 0
	total = a + b
	for _, v := range args {
		total += v
	}
	return total

}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...)
	}
	return 0

}
