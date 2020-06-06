package main

import "fmt"

func main() {
	/*
		多维数组
			一维素组：存储的多个数据是数值本身
			arr1 :=[3] int {1,2,3}

			二维数组: 存储的是一堆一维数组
			arr2 :=[3][4] int {{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	*/

	// 定义一个二维数组

	arr2 := [2][1]int{{1}, {2}}

	fmt.Println(arr2)
	fmt.Println(arr2[0][0])
	fmt.Println(len(arr2))

	//简短定义一个二维数组
	arr3 := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr3)
	fmt.Printf("二维数组的长度: %d\n", len(arr3))
	fmt.Printf("二维数组内的一维数组长度: %d\n", len(arr3[0]))

	fmt.Println("---------------------")

	//二维数组遍历

	//第一种 for循环遍历

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[j]); j++ {
			fmt.Print(arr3[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
	// 第二种 for range 遍历二维数组

	for _, arr := range arr3 {
		for _, value := range arr {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
