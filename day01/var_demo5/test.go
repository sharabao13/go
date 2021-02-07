package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var bytes []byte
	bytes = []byte{'a', 'b', 'c'}
	fmt.Println(bytes)
	// RuneCountInString 获取unicode 个数
	fmt.Println(utf8.RuneCountInString("我是谁"))

	//string转其他类型的
	if v, err := strconv.ParseBool("trufdsdfds"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	v, err := strconv.ParseInt("64", 16, 64)
	fmt.Println(v, err)

	heights := []int{6, 7, 9, 10, 5}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights); j++ {
			if heights[j] > heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
			}
		}
	}
}
