package main

import (
	"fmt"
)

func main() {
	seq := []string{"a", "b", "c", "d", "e"}

	index := 2
	//查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index:])

	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1])
	fmt.Println(seq)

	vsSlice := make([]int, 10)
	for i := 1; i < len(vsSlice); i++ {
		vsSlice[i] = i
	}
	fmt.Println(vsSlice)
	for k, v := range vsSlice {
		fmt.Println(k, v)
	}
}
