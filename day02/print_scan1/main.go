package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	   键盘输入练习
	*/
	var (
		a int
		b int
	)
	fmt.Println("请输入两个整数：")
	//fmt.Println("请输入两个整数: ")
	// 读取键盘数据，操作地址复制给变量
	fmt.Scanln(&a, &b)
	fmt.Printf("a的值: %d,b的值: %d\n", a, b)

	fmt.Println("请输入两个整数：")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("a: %d,b: %d\n", a, b)

	fmt.Println("请输入一个字符串: ")
	input := bufio.NewReader(os.Stdin)
	s1, _ := input.ReadString('\n')
	fmt.Printf("s1读到的数据: %s\n", s1)

}
