package main

import "fmt"

func main() {
	/*
		数组指针： 一个指针，一个数组地址
	       *[4]type
		指针数组： 首先是一个数组，存储的数据类型是指针
	  		[4]*type

			*[5]float 指针 一个存储了个5个浮点类型数据的数值指针
			*[3]string 指针 一个存储了个3个字符串类型数据的数值指针
			[3]*string	数组 存储了3个字符串指针地址的数组
			[5]*float64 数组 存储了5个浮点数据地址的数组
			*[3]*string 指针 存储了3个字符串指针地址的数组指针
			**[4]string 指针，存储了个4个字符串数组指针地址的指针
			**[4]*string 指针 存储了4个字符串的指针地址的素组，的指针的指针
	*/

	//创建一个数组
	arr1 := [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(&arr1)

	//创建一个数组，存储改数组的地址，数组指针
	var p1 *[3]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("p1值的地址：%p\n", p1)
	fmt.Printf("p1自身地址: %p\n", &p1)

	//操作指针改变数值元素的值
	(*p1)[0] = 20
	fmt.Println(arr1)
	p1[1] = 50 //简化方式
	fmt.Println(arr1)
	fmt.Println("-------------------")
	// 创建指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("---------------")

	//改变arr2数组的值，a 的值跟数组arr3的值不变
	arr2[0] = 100
	fmt.Println(arr2)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
	//fmt.Println(arr3)
	fmt.Println(a)
	fmt.Println("------------------")

	//改变指针数组的值，a的值会跟着改变，数组arr2的值不会变
	*arr3[1] = 200
	fmt.Println(arr2)
	fmt.Println(b)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}
