package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		function
	*/
	var a float64 = 28.11
	var b float64 = 20.22
	var ret float64

	ret = math.Max(a, b)
	fmt.Printf("最大值: %.2f\n", ret)

	testFun1 := isMax(12, 89)
	fmt.Println(testFun1)

	var c int = 21
	s := oddNum(c)
	fmt.Println(s)

}
func isMax(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}

func oddNum(c int) string {
	if c%2 == 1 {
		return "奇数"
	} else {
		return "偶数"
	}
}
