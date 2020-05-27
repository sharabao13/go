package main

import "fmt"

func main() {
	/*
			2-100的素数
			素数只能被1 和自身整除的数
		   2,3，5，  7

	*/
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
