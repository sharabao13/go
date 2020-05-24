package main

import "fmt"

func main() {
	/*
			switch 分支
				语法：
					switch 变量名{
						case 数值：分支1
						case 数值：分支2
						case 数值：分支3
						..
						default:
							最后一个分支
					}
			注意事项 ：
				1.switch 可以作用在其他数据类型上，case 后的数值必须跟switch作用的变量类型一致
		        2.case是无序的
				3.case值是唯一的
	*/
	Qua := 5
	switch Qua {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("无效数据")
	}
	fmt.Println("main...over...")
}
