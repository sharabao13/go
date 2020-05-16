// 变量声明类型
package main

import "fmt"

func main() {

	//声明变量并赋值
	var id int = 1
	var stdName string = "bao"

	//变量类型推导，根据变量的值判断变量的类型 格式 var 变量名称 =
	var id1 = 1
	var stdName1 = "bao1"

	// 简短变量声明  格式 变量名称 :=
	id2 := 1
	stdName2 := "bao2"

	// 同一作用域变量名称不能一样
	fmt.Println(id)
	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(stdName)
	fmt.Println(stdName1)
	fmt.Println(stdName2)
}
