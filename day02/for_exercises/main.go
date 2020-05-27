package main

import (
	"fmt"
	"math"
)

func main() {
	/*
			水仙花数 [100,999]
			自身位数的立方和等于自身，
			153=1三次方+5三次方+3三次方
		    math.Pow(x,y) y代表几次幂 float64类型
	*/
	for i := 100; i < 1000; i++ {
		x := i / 100     //百位数
		y := i / 10 % 10 //十位数
		z := i % 10
		// math.Pow求几次幂的方法
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}

	// 另一种写法 三位数
	// 百位数 1-9
	// 十位数 0-9
	// 个位数 0-9

	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Println(n)
				}
			}
		}
	}
}
