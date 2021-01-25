package main

import (
	"errors"
	"fmt"
	"strings"
)

// 高阶函数  函数作为参数 回调函数

//定义加法函数
func add(x, y int) int {
	return x + y
}

//定义减法函数
func sup(x, y int) int {
	return x - y
}

//定义乘法函数
func mul(x, y int) int {
	return x * y
}

//定义除法函数
func div(x, y int) int {
	if y == 0 {
		fmt.Println("被除数不能为0")
		return 0
	}
	return x / y
}

//定义操作加减函数
func calc(x, y int, oper func(int, int) int) int {
	return oper(x, y)
}

//定义四则运算函数
func calc1(op string) (func(int, int) int, error) {
	switch op {
	case "+":
		return add, nil
	case "-":
		return sup, nil
	case "*":
		return mul, nil
	case "/":
		return div, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err

	}

}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	//calc 函数调用加法
	add1 := calc(1, 2, add)
	fmt.Println(add1)
	//calc 函数调用减法
	sup1 := calc(9, 7, sup)
	fmt.Println(sup1)
	//调用四则运算

	fun1, err := calc1("+")
	if err != nil {
		fmt.Println(err)
	}
	res1 := fun1(10, 9)
	fmt.Println(res1)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
