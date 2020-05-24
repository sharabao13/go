// if 条件判断
package main

import "fmt"

func main() {
	score := 59
	if score >= 60 {
		if score < 100 {
			if score >= 90 {
				fmt.Println("成绩优秀")
			} else if score >= 80 && score < 90 {
				fmt.Println("成绩良好")
			} else if score >= 60 && score < 80 {
				fmt.Println("成绩合格")
			}
		}
	} else {
		fmt.Println("成绩不合格")
	}
	// if 定义内部作用域的变量，外部不生效
	if age := 18; age >= 18 {
		fmt.Println("yong")
	} else if age < 18 {
		fmt.Println("small")
	}

}
