package main

import "fmt"

func main() {
	/*
		defer 延迟，推迟
		在go语言中 defer 用来对一个函数或一个方法的延迟执行

		1.defer函数或方法，一个函数或方法被延迟执行
		2.defer用法：
			1.对象.close(),临时文件删除
				文件.open()
				defer .close()
				读或写

			2.go语言中关于异常的处理，使用panic()或recover（）
				panic()函数用于引发恐慌，导致程序中断执行
				recover()函数用于程序恢复，recover()语法上要求必须在defer中执行

		3.如何多个defer： 先延迟的后执行，后延迟的先执行

		4.defer函数传递参数： 在函数传递的时候，参数已经传递，只是代码暂时延迟执行

		5.defer函数注意点
			1.当外网函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外网函数才会真正的被执行完毕
			2.当执行外网函数中的return语句时，只有其中所有的延迟函数都执行完毕，外网函数才会真正返回
			3.当外网函数中的代码引发恐慌时，只有其中所有的延迟函数都执行完毕，该运行时的恐慌才会真正被扩展至调用函数
	*/
	//fmt.Println("Hello")
	//fmt.Println("123")
	//defer fun1("World") //函数被延迟执行，放在外围程序执行后执行
	//fmt.Println("789")
	//defer fun1("Over")

	a := 2
	fmt.Println("main中a的初始值", a) //2
	defer fun2(a)                //2 别延迟执行 此时a传递参数2,只是延迟执行
	a++                          //3
	fmt.Println("main中a的最后值", a) //3
	fmt.Println(fun3())
}
func fun1(s string) {
	fmt.Println(s)
}

func fun2(a int) {
	fmt.Println("函数中a的值", a)
}

func fun3() int {
	fmt.Println("fun3中函数真正被执行")
	defer fun1("delete")
	return 0
}
