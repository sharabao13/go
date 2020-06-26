package main

import "fmt"

func main() {
	/*
		面向对象世界中的接口的一般定义是“接口定义对象的行为”。它表示让指定对象应该做什么。实现这种行为的方法(实现细节)是针对对象的。

		在Go中，接口是一组方法签名。当类型为接口中的所有方法提供定义时，它被称为实现接口。它与OOP非常相似。接口指定了类型应该具有的方法，
		类型决定了如何实现这些方法。

		它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

		接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了该接口。

		go语言中，接口和类型的实现关系，是非侵入式

		1.当需要接口类型对象时，可以使用任意实现类对象代替
		2.接口对象不能访问，实现类中的熟悉
	*/

	// 创建mouse对象
	m1 := Mouse{name: "dell"}
	fmt.Println(m1)

	//创建U盘对象
	u1 := FlashDisk{name: "闪迪64G"}
	fmt.Println(u1)

	//调用测试方法
	testInterface(m1)

	testInterface(u1)

	//当需要接口类型对象时，可以使用任意实现类代替
	//接口对象不能访问实现中的对象

	var usb USB
	usb = m1
	fmt.Println(usb)
	testInterface(usb)

	u1.delete()

	//接口的用法
	//1.一个函数如果接收参数为接口类型，那么实际上可以传入改接口类的任意实现类型对象作为参数
	//2.定义一个类型为接口类型，实际上可以赋值任意实现类对象

	var arr3 [3]USB
	arr3[0] = m1
	arr3[1] = u1
	fmt.Println(arr3)

	testInterface(arr3[0])
}

//定义接口
type USB interface {
	start() //usb开始工作
	end()   //USB结束工作
}

// 定义一个鼠标实现类
type Mouse struct {
	name string
}

//定义一个u盘实现
type FlashDisk struct {
	name string
}

//定义鼠标实现方法
func (m Mouse) start() {
	fmt.Println(m.name, "鼠标开始工作")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束工作")
}

//定义U盘实现方法
func (f FlashDisk) start() {
	fmt.Println(f.name, "U盘开始工作")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "U盘退出工作")
}

func (f FlashDisk) delete() {
	fmt.Println(f.name, "删除数据")
}

//定义一个测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

///* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//	method_name2 [return_type]
//	method_name3 [return_type]
//	...
//	method_namen [return_type]
//}
//
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
//	/* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
//	/* 方法实现*/
//}
