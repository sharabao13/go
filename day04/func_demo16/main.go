package main

import "fmt"

// 函数变量 将一个函数赋值给一个变量
func testFun(a, b int, n ...string) {
	fmt.Printf("输入%d输出%d\n", a, a)
	fmt.Printf("输入%d输出%d\n", b, b)
	fmt.Printf("输入%s输出%s\n", n, n)

}

func getScore(s int) {
	if s >= 60 {
		if s < 100 {
			if s >= 90 {
				fmt.Printf("成绩优秀")
			} else if s >= 80 && s < 90 {
				fmt.Printf("成绩良好")
			} else if s >= 60 && s < 80 {
				fmt.Printf("成绩合格")
			}
		}
	} else {
		fmt.Printf("成绩不合格")
	}
}

func oddNumber(n int) {
	if n%2 != 0 {
		fmt.Printf("%d是奇数", n)
	} else {
		fmt.Printf("%d是偶数", n)
	}
}

func getSum(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("0-%d的总和是: %d\n", n, sum)
}

func main() {
	//oddNumber(11)

	getSum(100)

	getScore(56)

	testFun(1, 2, "test", "test1", "test3")

}
