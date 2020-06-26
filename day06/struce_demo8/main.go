package main

import "fmt"

//创建猫的结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	/*
		张老太养了2只猫猫，一只名叫小白，今年三岁，白色，还有一只叫小花，今年100岁，花色
		请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色，如果用户输入的小猫名字错误，则显示，没有这只猫猫
	*/

	//1. 使用变量处理

	//var cat1Name string = "小白"
	//var cat1Age int = 3
	//var cat1Coller string  ="白色"
	//
	//var cat2Name string = "小美"
	//var cat2Age int = 100
	//var cat2Coller string  ="花色"
	//
	////2. 使用数组方式
	//var catNames [2]string = [...]string{"小白","小花"}
	//var catAges [2]int = [...]int{3,100}
	//var catCollers [2]string = [...]string{"白色","花色"}

	/*现有技术解决的缺点分析
		1.不利于数据的管理和维护，因为名字，年龄，颜色的属性都属于一直猫
	  	2.如果我们希望对一只猫的属性{名字，年龄，颜色}进程操作（绑定方法）也不好处理
		3.使用结构体新技术
	*/

	// 3.使用结构体方式
	//创建一个变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("猫的信息如下：")
	fmt.Println("name: ", cat1.Name)
	fmt.Println("age: ", cat1.Age)
	fmt.Println("color: ", cat1.Color)

}
