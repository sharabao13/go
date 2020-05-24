package main

import "fmt"

func main() {
	/*
		switch 分支
			1.break: 可以使用在switch中，也可以使用在for循环中
					强制结束case语句，结束switch分支
			2.fallthrough 贯穿满足case的后续的紧邻的case
				如需贯穿多个case可在case最后一行添加fallthrough
	*/
	num := 2
	switch num {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2: //打印三行
		fmt.Println("我是熊二")
		break //强制结束case跳出switch分支，此时只打印一行
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
	}
	fmt.Println("main...over...")
	fmt.Println("---------------------------------------")
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough //贯穿3的case case3的语句也会执行
	case 3:
		fmt.Println("第三季度")
		fallthrough //贯穿4的case case4的语句也会执行
	case 4:
		fmt.Println("第四季度")
	}
	fmt.Println("main...over...")
}
