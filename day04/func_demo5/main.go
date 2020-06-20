package main

import "fmt"

func main() {
	/*
		  函数参数传递
			1.值传递，传递的是数据的副本，修改数据对于原始数据不会改变
				值类型的数据默认都是值传递，基础类型，array struct

			2.引用传递，传递的数据地址，修改数据会对原始数据进行修改
				引用类型的数据都是引用传递，slice，map，channel

	*/

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("函数调用前数组的数据：", arr1)

	getArry(arr1)
	//值传递类型，修改数据过程中对于原始数据不会变化
	fmt.Println("函数调用后数组的数据：", arr1)
	fmt.Println("-----------------------------")

	slice1 := []int{7, 8, 9}
	fmt.Println("函数调用前切片的原始数据：", slice1)
	getSlice(slice1)
	//引用类型数据中间修改数据，原始数据也会随着改变
	fmt.Println("函数调用后切片的原始数据：", slice1)
}

//创建一个数组函数

func getArry(arr2 [5]int) {
	fmt.Println("函数中，数组的数据: ", arr2)
	arr2[0] = 88
	fmt.Println("函数中，修改数据后的数据：", arr2)
}

//创建一个切片函数
func getSlice(slice2 []int) {
	fmt.Println("函数中，切片的原始数据：", slice2)
	slice2[0] = 10
	fmt.Println("函数中，切片修改后的原始数据：", slice2)
}
