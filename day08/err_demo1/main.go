package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		错误是什么?

		错误指的是可能出现问题的地方出现了问题。比如打开一个文件时失败，这种情况在人们的意料之中 。

		而异常指的是不应该出现问题的地方出现了问题。比如引用了空指针，这种情况在人们的意料之外。可见，错误是业务过程的一部分，而异常不是 。

		Go中的错误也是一种类型。错误用内置的error 类型表示。就像其他类型的，如int，float64，。错误值可以存储在变量中，从函数中返回，等等。
	*/

	f, err := os.Open("test.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("err")
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("op", ins.Op)
			fmt.Println("path", ins.Path)
			fmt.Println("err", ins.Err)
		}
		return

	}
	fmt.Println(f.Name(), "打开文件成功")
}
