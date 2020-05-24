// for 循环
package main

import "fmt"

// for 基本格式
func main() {
	// 基本格式 变量作用域内使用
	//for i := 0;i <10;i++{
	//	fmt.Println(i)
	//}

	// 变量在外部的使用
	//i := 5
	//for ;i < 10;i++{
	//	fmt.Println(i)
	//}

	//初始语句跟结束语句省略
	//i := 5
	//for ;i < 10;{
	//	fmt.Println(i)
	//	i++
	//}
	for i := 1; i <= 2; i++ {
		fmt.Println(i)
	}
}
