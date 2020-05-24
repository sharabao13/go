package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		输入/输出
		输出：
			print 打印
			printf 格式化打印
			println 打印之后换行
		格式化打印占位符
			%T  数据类型
			%t  布尔类型
			%v  原样输出
			%#v 带数据类型的原样输出
			%b  二进制输出
			%d  十进制
			%o  八进制
			%x  十六进制 0-9 a-f
			%X  十六进制 0-9 A-F
			%f  浮点型
			%c  打印字符
			%p  内存地址

	*/
	a := 1
	s := "str"
	s1 := 'a'
	fmt.Println("-----------------")
	fmt.Printf("%v,%d\n", a, a)
	fmt.Printf("%#v,%s\n", s, s)
	fmt.Printf("%#c,%s\n", s1, s1)
	fmt.Println("-----------------")

	/*
		从键盘读取数据
		Scanln(&x,&y) //& 从操作地址来赋值给变量
		Scanf("%d,%f",&x,&t) 格式化读取键盘的值赋值给变量
		bufio 调用NewReader(os.stdin) 方法
	*/

	/*
		var x int
		var y float64
		fmt.Println("请输入一个整数，一个浮点数：")
		// 读取键盘的输入，通过操作地址来赋值给x,y变量 阻塞式方法
		fmt.Scanln(&x, &y)
		fmt.Printf("x的数值：%d,y的数值: %.2f\n", x, y)

		// scanf 输入的值需安装format内的格式 如%d，%d 要输入 1,1 按符号隔开
		fmt.Scanf("%d %f", &x, &y)
		fmt.Printf("x:%d,y:%.2f\n", x, y)
	*/
	fmt.Println("请输入一个字符串：")
	// bufio NewReader(os.Stdin)  // 标准输入
	reader := bufio.NewReader(os.Stdin)
	// 分隔
	s3, _ := reader.ReadString('\n')
	fmt.Printf("读到的数据: %s\n", s3)
}
