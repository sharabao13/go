package main

import "fmt"

func main() {
	/*
		for range 练习遍历数组
	*/
	// for range 遍历一维数组

	arr1 := [...]string{"apple", "banana", "pear"}
	for _, val := range arr1 {
		fmt.Print(val, "\t")
	}
	fmt.Println()
	fmt.Println("-------------------")

	arr2 := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, arr := range arr2 {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}
