package main

import "fmt"

func main() {
	/*
		冒泡排序
	*/
	// 练习 45,3,66,33,12,9,22,87,1从大到小排序

	arr1 := [...]int{45, 3, 66, 33, 12, 9, 22, 87, 1}
	//for i := 0; i < 8; i++ {
	//	if arr1[i] > arr1[i+1] {
	//		arr1[i], arr1[i+1] = arr1[i+1], arr1[i]
	//	}
	//}
	//fmt.Println(arr1)

	for j := 0; j < len(arr1); j++ {
		for i := 0; i < len(arr1)-1; i++ {
			if arr1[i] > arr1[i+1] {
				arr1[i], arr1[i+1] = arr1[i+1], arr1[i]
			}
		}
	}
	fmt.Println(arr1)
}
