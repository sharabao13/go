package main

import "fmt"

// 数据元素遍历
func main() {
	var arr1 [4]int
	var arr2 [4]string
	var arr3 [2]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)
}
