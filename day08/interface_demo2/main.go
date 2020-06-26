package main

import "fmt"

func main() {
	/*
		定义接口练习
	*/

	var phone phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}

//定义一个phone的接口类型
type phone interface {
	call()
}

//定义一个nokiaPhone 的结构体
type NokiaPhone struct {
}

//定义一个NokiaPhone 的实现方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia i can call you ")
}

//定义一个Iphone的结构体
type Iphone struct {
}

//定义一个Iphone 的实现方法
func (iphone Iphone) call() {
	fmt.Println("I am Iphone i can call you ")
}
