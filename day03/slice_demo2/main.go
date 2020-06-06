package main

import "fmt"

func main() {
	/*
	  切片练习
	*/

	// 定义一个切片

	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	//make 定义一个切片  //长度为2容量为5
	s2 := make([]int, 2, 5)
	fmt.Println(s2)

	// 向s2切片追加内容
	s2 = append(s2, 0, 1)
	fmt.Println(s2)

	s2[0] = 1
	fmt.Println(s2)

	s4 := []int{7, 8, 9}

	s2 = append(s2, s4...)
	fmt.Println(s2)
	// 切片遍历
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], "\t")
	}
	fmt.Println()

	// for range 切片
	for _, val := range s2 {
		fmt.Print(val, "\t")
	}
	fmt.Println()
}
