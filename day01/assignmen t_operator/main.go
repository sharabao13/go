package main

import "fmt"

func main() {
	/*
		赋值运算符
		= += -= *= /= %= <<= >>= &= |= ^=..
		将右边的值进行运算在赋值给左边的
		+= a += b  相当于 a = a + b
	*/
	a := 3
	a += 3
	fmt.Println(a)
}
