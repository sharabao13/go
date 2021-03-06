package main

import "fmt"

func main() {
	/*
				   1.for循环标准写法
						for 表达式1;表达式2;表达式3 {
								循环体
							}
				   2.同时省略表达式1,表达式3
			        	for 表达式2 {  //分号可去除
								循环体
							}
		 		   3.同时省略三个表达式
						for {

							}
				    //相当于while true
					注意点: for循环中省略了表达式2,相当于条件为true
				   4.其他写法
						省略表达式1: 变量要在循环体外面声明
						省略表达式2: 循环永远成立,死循环
						省略表达式3: 变量变化写在循环体内


	*/

	// 1. for 循环标准写法
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------------")

	//2. for 循环省略表达式1 表达式3
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
