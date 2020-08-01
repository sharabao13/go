package main

import "fmt"

func main() {
	/*
		panic 恐慌
		recover "恢复"
		go 语言中用panic() recover() 实现程序中极特殊的异常处理
		panic() 当前程序进入恐慌，终端程序运行
		recover() 让程序恢复，必须在defer函数中执行
	*/
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("程序恢复", msg)
		}
	}()
	Funa()
	defer myprints("defer main .....3")
	funb()
	defer myprints("defer main .....4")
	fmt.Println("main....over")
}

//定义打印函数
func myprints(s string) {
	fmt.Println(s)
}

// 定义函数A
func Funa() {
	fmt.Println("我是函数A")
}
func funb() { //外围函数
	fmt.Println("我是函数B")
	defer myprints("defer funB .....1")
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
		if i == 5 {
			panic("funB 恐慌了")
		}
	}
	// 当外围函数发生panic时已经defer的函数会被执行，当全部defer执行完毕后panic会执行
	defer myprints("defer funB .....2")
}
