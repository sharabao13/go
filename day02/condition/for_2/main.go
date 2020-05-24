// 九九乘法表
package main

import "fmt"

func main() {
	// 打印层数
	for i := 1; i <= 9; i++ {
		// 打印 行数
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}
