package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}
func main() {
	//错误处理
	v1, err := div(2, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v1)

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name(), "打开文件成功")
}
