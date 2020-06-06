package main

import "fmt"

func main() {
	/*
		冒泡排序
	*/

	arr1 := [...]int{1, 3, 6, 8, 92, 11, 3, 0, 65, 9}
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-1; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
		fmt.Println(arr1)
	}
	fmt.Println(arr1)
}
