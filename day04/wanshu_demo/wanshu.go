package main

import (
	"fmt"
)

func calc(n int) bool {
	var sum int

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n == sum
}

func wanshu(n int) {
	for j := 1; j <= n; j++ {
		if calc(j) {
			fmt.Println(j)
		}
	}
}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	wanshu(a)
}
