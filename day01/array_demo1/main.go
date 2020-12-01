package main

import "fmt"

// 数据元素遍历
func main() {
	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 100
		fmt.Println(n[i])
	}

	for j := 0; j < 10; j++ {
		fmt.Println(n[j])
	}
}
