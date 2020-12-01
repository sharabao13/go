package main

import "fmt"

//  遍历数据
func main() {
	a := [...]int{1, 2, 3, 40}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
