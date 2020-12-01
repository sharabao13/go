package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 6)
	slice1[1] = 8
	slice1[2] = 9
	slice1[3] = 10
	slice1[4] = 11
	fmt.Println(slice1)

	for _, k := range slice1 {
		fmt.Println(k)
	}
	slice2 := [5]int{1, 2, 3, 4, 5}
	var slice3 []int = slice2[0:3]
	for _, k1 := range slice2 {
		fmt.Println(k1)
	}
	fmt.Println(slice2)
	for _, k2 := range slice3 {
		fmt.Println(k2)
	}
	fmt.Println(slice3)
	fmt.Println("--------------------------")

	slice4 := [5]int{78, 100, 200}
	fmt.Println(slice4)
	slice5 := slice4[:]
	slice6 := slice4[:]
	slice5[0] = 100
	fmt.Println("修改赋值后的值", slice4)
	slice6[2] = 1000
	fmt.Println("修改赋值后的值", slice4)

	var numbers []int
	if numbers == nil {
		fmt.Println("切片是空的")
	}

}
