package main

import "fmt"

func main() {
	/*
		函数指针
			一个指针，指向了函数的地址
			go语言中 function默认看作一个指针，没有*
				slice map function
		指针函数
			一个函数，该函数的返回值是指针
	*/
	var a func()
	a = fun1
	a()
	fmt.Println("----------------")

	arr1 := fun2()
	fmt.Printf("arr1的类型%T,地址：%p,数值： %v\n", arr1, &arr1, arr1)
	fmt.Println("----------------")

	arr2 := fun3()
	fmt.Printf("arr2的类型: %T，地址：%p,数值：%v\n", arr2, &arr2, arr2)
	fmt.Println("----------------")

}

func fun1() {
	fmt.Println("fun()...")
}

func fun2() [4]int { //普通函数
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	return &arr
}
