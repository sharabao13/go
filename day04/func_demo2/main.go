package main

import "fmt"

func main() {
	/*
		函数的参数
			1.1 参数的使用
				形式参数：函数定义时用于接收外部传入的参数，形参

				实际参数：传给形参的实际数据，实参

			1.2 函数调用
				函数名必须匹配
				形参跟实参需一一对应，顺序，个数，类型
	*/

	//函数参数的初步使用
	getSum(10)
	fmt.Println("---------------------")

	getSum(20)
	fmt.Println("---------------------")

	getSum(100)
	fmt.Println("---------------------")

	getSum(1000)
	fmt.Println("---------------------")

	multiN(2)
	fmt.Println("---------------------")

	multiN(6)
	fmt.Println("---------------------")

	multiN(9)
	fmt.Println("---------------------")

	getExerciese(20)
	fmt.Println("---------------------")

}

//创建1-n的和的函数
//（）n 形参，加类型
func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是: %d\n", n, sum)
}

//创建一个n*n乘法表
func multiN(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

}

//创建一个2-n之内的素数

func getExerciese(n int) {
	fmt.Printf("打印2-%d的素数\n", n)
	for i := 2; i <= n; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
