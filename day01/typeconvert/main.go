// 数据类型转换
package main

import "fmt"

func main() {
	/*
			语法格式 Tyep(value)
			常数： 在有需要的时候，会自动转换
		    变量： 需手动转换
			注意项
				兼容类型可以转换
	*/
	var a1 int8
	a1 = 10

	var a2 int16
	//a2 = a1  //cannot use a1 (type int8) as type int16 in assignment
	// 强制类型转换
	a2 = int16(a1)
	fmt.Println(a1, a2)

	a3 := 3.14
	var a4 int8
	// 数据类型转换
	a4 = int8(a3)
	fmt.Println(a4)

	a5 := 10
	a3 = float64(a5)
	fmt.Println(a5, a3)

	//不同的数据类型无法强制转换
	//	b1 := false
	//	a5 = int(b1)  //cannot convert b1 (type bool) to type int
	//	fmt.Println(b1, a5)
}
