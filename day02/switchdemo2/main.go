package main

import "fmt"

func main() {
	/*
			switch {  //变量为空的时候，switch默认值是true 布尔类型
		    case true 分支
			case false 分支
			}
			case 后可跟多个数值
			case 1,3,5,7,9:

			switch 可以多一条初始化语句
				switch 初始化语句;var{

					}
	*/

	/*
		switch 成绩分支练习
	*/

	var Score int
	fmt.Println("请输入一个分数： ")
	fmt.Scanln(&Score)
	switch {
	case Score < 60:
		fmt.Printf("成绩不及格\n")
	case Score >= 60 && Score < 70:
		fmt.Printf("成绩及格\n")
	case Score >= 70 && Score < 80:
		fmt.Printf("成绩中等\n")
	case Score >= 80 && Score < 90:
		fmt.Printf("成绩良好\n")
	case Score >= 90 && Score < 100:
		fmt.Printf("成绩优秀\n")
	default:
		fmt.Printf("请输入0-100的成绩\n")
	}
	fmt.Println("main...over...")
}
