package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		error 内置的数据类型 数据接口
		定义方法 error() string

		使用go语言提供好的包
			errors包下的函数: New(),创建一个error对象

			fmt包下的errorf
				func Errorf(format string, a ...interface{}) error
	*/

	//创建一个error对象
	err1 := errors.New("测试")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	//创建另一个error方法
	err2 := fmt.Errorf("错误码: %d\n", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	a1 := checkAge(-30)
	if a1 != nil {
		fmt.Println(a1)
		return
	}
	fmt.Println("程序正常")
}

// 创建一个函数，验证年龄是否合法，如果为负数，就返回error
func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		err := fmt.Errorf("给定的年龄不合法: %d\n", age)
		return err
	}
	fmt.Println("年龄合法")
	return nil
}
