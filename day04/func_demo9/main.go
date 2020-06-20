package main

import "fmt"

func main() {
	/*
		递归函数 一个函数自己调用自己，
			一个函数有一个出口，逐渐向出口靠近

	*/

	sum := getSum(2)
	fmt.Println(sum)

	//阶乘调用
	var i int = 4
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	//Fibonacci数列
	n := 2
	fmt.Printf("%d的Fibonacci数是：%d", n, Fibonacii(n))
}

// 1-5的和的递归函数
func getSum(n int) int {
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}

//阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

//fibonacci 数列
/*
	1 2 3 4 5 6 7
	1 1 2 3 5 8 13
*/

func Fibonacii(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacii(n-1) + Fibonacii(n-2)
}
