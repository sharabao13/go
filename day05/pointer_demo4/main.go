package main

import "fmt"

func main() {
	/*

	 */
	a := 10
	fmt.Println("函数调用前a的值:", a)
	fun1(a)
	fmt.Println("函数调用后a的值:", a)
	fmt.Println("----------------")

	fun2(&a)
	fmt.Println("函数调用后a的值:", a)
	fmt.Println("----------------")

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前arr1的值: ", arr1)
	fun3(arr1)
	fmt.Println("函数调用后arr1的值: ", arr1)
	fmt.Println("----------------")

	fun4(&arr1)
	fmt.Println("函数调用后a的值:", arr1)
	fmt.Println("----------------")
}

func fun1(num int) { //值传递
	fmt.Println("func()中num的值：", num)
	num = 100
	fmt.Println("func()修改num的值: ", num)
}

func fun2(p1 *int) { //引用传递
	fmt.Println("func()中num的值：", *p1)
	*p1 = 100
	fmt.Println("func()修改num的值: ", *p1)

}

func fun3(arr [4]int) { //值传递
	fmt.Println("fun()中arr的值: ", arr)
	arr[0] = 100
	fmt.Println("fun()修改后的值: ", arr)
}

func fun4(p1 *[4]int) {
	fmt.Println("fun()中arr的值: ", p1)
	p1[0] = 100
	fmt.Println("fun()修改后的值: ", p1)
}
