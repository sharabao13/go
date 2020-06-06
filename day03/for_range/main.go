package main

import "fmt"

func main() {
	/*
		数组的遍历
			1.依次访问数组中的元素
				arr[0],arr[1],arr[2],arr[3],arr[n]

			2.for 循环遍历数组中的内容
				for i:=0;i<len(arr);i++{
						arr[i]
					}

			3.for range 方法遍历数组内容
				不需要操作数组的下标，到达数组的末尾，结束for range 循环
				每次循环从数组中获取下标和对应的值
				语法 for index,value := range arr{
					}
	*/

	// 第一种数组遍历
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])

	fmt.Println("--------------------")
	// for循环遍历数组

	arr2 := [...]string{"apple", "banana", "pear"}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	fmt.Println("---------------------")
	//for range 遍历数组中的数据
	for index, value := range arr2 {
		fmt.Printf("下标是: %d 值是: %s\n", index, value)
	}
	fmt.Println("---------------------")

	//for range 不打印下标
	for _, v := range arr2 {
		fmt.Println(v)
	}
}
