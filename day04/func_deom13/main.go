package main

import "fmt"

func main() {
	/*
			go语言支持函数式编程
			支持将一个函数作为另一个函数的参数
			也支持将另一个函数作为函数的返回值
		闭包：
			一个外层函数中，有内层函数，该内城函数中，会操作外层函数的局部变量（外层函数中或内层函数指定的变量），改内层函数的返回值就是这个内层函数
			这个外层函数或内层函数的局部变量，统称为闭包结构

			局部变量的生命周期将会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁
			但是，闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内存函数还要继续调用
	*/

	//闭包结构函数调用
	res1 := increment()
	fmt.Printf("%T\n", res1) //打印函数类型
	fmt.Println(res1)        //打印函数内存地址
	v1 := res1()
	fmt.Println(v1)
	v2 := res1()
	fmt.Println(v2)
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
}

func increment() func() int { //外层函数
	//定义了一个局部变量
	i := 0
	//定义了一个匿名函数，返回值 自增
	fun := func() int { //内层函数
		i++
		return i
	}
	//返回匿名函数
	return fun
}
