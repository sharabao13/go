package main

import "fmt"

//go 语言中推荐使用驼峰式命名
//单独声明变量

//var name string
//var age int
//var isok bool

//批量声明
var (
	name string //""
	age  int    //0
	isok bool   //false
)

func main() {
	name = "tae"
	age = 18
	isok = true
	//变量赋值必须使用，不使用不能编译
	fmt.Print(isok) //直接打印
	fmt.Println()
	fmt.Printf("name:%s\n", name) //格式化输出 %s占位符
	fmt.Println(age)              //打印完换行
}
