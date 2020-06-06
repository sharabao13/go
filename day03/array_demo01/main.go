package main

import "fmt"

func main() {
	/*
		数据类型：
			1.基本类型：整数、浮点数、布尔、字符串
			2.符合类型：array slice,map,sturce,pointer,function,channel
		数组
			概念： 存储一组形同数据类型的结构
			语法：
					1. var [] 数据类型
					2.定义数组的同时将数组同时赋值
						var [] 数据类型{值，值，。。。}
			通过下标访问：arr[]
			长度和容量 len cap

	*/
	// 数组创建
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	//访问数组全部内容
	fmt.Println(arr1)

	//访问数组索引位置
	fmt.Println(arr1[2])

	//定义的属组大小不可变，值可变
	arr1[2] = 100
	fmt.Println(arr1[2])
	fmt.Println(arr1)

	//获取属组的长度及容量
	// len()方法 长度 也可在slice map string..使用
	// cap()方法 容量 也可在slice map string..使用
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	//定义属组并同时赋值
	var arr2 = [3]string{"apple", "banana", "bear"}
	fmt.Println(arr2)

	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3)

	//定义数组并同时给下标赋值，没赋值部分默认空值
	var arr4 = [3]int{0: 1}
	fmt.Println(arr4)

	//简短定义数组[...]

	a := [...]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{2: "apple"}
	fmt.Println(b)

}
