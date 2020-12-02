package main

import "fmt"

func main() {
	//对维切片
	var slice [][]int
	slice = [][]int{{10}, {20, 30}}
	fmt.Println(slice)

	//二维数组添加元素
	slice[0] = append(slice[0], 100)
	fmt.Println(slice)
}
