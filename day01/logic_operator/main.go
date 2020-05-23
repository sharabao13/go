package main

import "fmt"

func main() {
	/*
			逻辑运算符，返回值布尔类型
				逻辑与 && 条件全为真，返回真，有个条件为假，返回假
		         & 与 短路与 && 执行语句遇到一个条件为假，则返回结果为假
				逻辑或 || 条件全假 ，返回假，有个条件为真，返回真
				| 与 短路与 || 执行语句遇到一个条件为真，则返回结果为真
				逻辑非 ！ 取反 条件假 返回真，条件真返回假
	*/
	f1 := true
	f2 := false
	f3 := true
	f4 := false
	res1 := f1 && f2
	fmt.Printf("%t\n", res1) //f

	res2 := f1 && f2 && f3
	fmt.Printf("%t\n", res2) //f

	res3 := f1 || f2
	fmt.Printf("%t\n", res3) //t
	res4 := f2 || f4
	fmt.Printf("%t\n", res4) //f

	res5 := f1 || f2 || f3 || f4
	fmt.Printf("%t\n", res5) //t

	res6 := !f1
	fmt.Printf("%t\n", res6) //f

	res7 := !f2
	fmt.Printf("%t\n", res7) //t

	a := 8
	b := 6
	c := 10
	res8 := a > b && c%a == b && a < (c/b)
	fmt.Printf("%t\n", res8) // f

	res9 := b*2 < c || a/b != 0 || c/a > b //t
	fmt.Printf("%t\n", res9)

	res10 := !(c/a == b)
	fmt.Printf("%t\n", res10) // t
}
