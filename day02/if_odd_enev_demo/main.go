package main

import "fmt"

func main() {
	/*
	  if  判断奇数 偶数
	*/
	var a int
	fmt.Println("请输入一个整数： ")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Printf("您输入的数是: %d,%d是一个偶数\n", a, a)
	} else {
		fmt.Printf("您输入的数是: %d,%d是一个奇数数\n", a, a)
	}
}
