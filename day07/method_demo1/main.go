package main

import "fmt"

func main() {
	/*
		Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集

		方法只是一个函数，它带有一个特殊的接收器类型，它是在func关键字和方法名之间编写的。接收器可以是struct类型或非struct类型。接收方可以在方法内部访问。

		方法能给用户自定义的类型添加新的行为。它和函数的区别在于方法有一个接收者，给一个函数添加一个接收者，那么它就变成了方法。接收者可以是值接收者，也可以是指针接收者。

		在调用方法的时候，值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。

		也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

		方法 一个方法就是包含了一个接收者的函数，接收者可是命名类型或者结构体类型的一个值或是一个指针
			所有给定类型的方法属于该类型的方法集

		语法
			func (接收者) 方法名（参数列表）(返回值列表){}

		总结 method 同函数类似，区别需要有接受者

		对比函数
		 1. 意义
			方法： 某个类别的行为功能，需要指定的接收者使用
			函数： 一段独立功能的代码，可以直接调用

		2. 语法
			方法： 方法名可以相同，只需接收者不同
			函数： 命名不能冲突
	*/

	//创建一个工人的对象
	w1 := Worker{name: "a", age: 18, gender: "male"}
	fmt.Println(w1)

	//调用work行为方法
	w1.work()

	w2 := &Worker{name: "b", age: 20, gender: "fmale"}
	fmt.Printf("%T\n", w2)
	w2.work()

	//调用rest行为方法指针
	w2.rest()

	w2.printInfo()
	w3 := Cat{name: "小白", age: 3, color: "白色"}
	w3.printInfo()
}

// 创建一个工人的结构体
type Worker struct {
	name   string
	age    int
	gender string
}

type Cat struct {
	name  string
	age   int
	color string
}

//定义一个行为方法
func (w Worker) work() {
	fmt.Println(w.name, "在工作...")
}

//定义一个指针的行为方法

func (p *Worker) rest() {
	fmt.Println(p.name, "在休息")
}

func (p *Worker) printInfo() {
	fmt.Printf("姓名: %s,年龄: %d,性别: %s\n", p.name, p.age, p.gender)
}

func (p *Cat) printInfo() {
	fmt.Printf("姓名: %s,年龄: %d,颜色: %s\n", p.name, p.age, p.color)
}
