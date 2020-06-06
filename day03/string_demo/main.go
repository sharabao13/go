package main

import "fmt"

func main() {
	/*
	 字符串
	*/

	nameSpace := "hello world"
	fmt.Println(nameSpace)

	// 遍历字符串
	for i := 0; i < len(nameSpace); i++ {
		//fmt.Println(nameSpace)
		//fmt.Print("\t", nameSpace[i])
		//按字节格式输出
		//fmt.Println(nameSpace[i])

		//按字符格式输出 %c
		fmt.Printf("\t%c\n", nameSpace[i])
	}
	fmt.Println("---------------------")

	//for range 遍历字符串

	for _, v1 := range nameSpace {
		fmt.Printf("%c\n", v1)
	}
	fmt.Println("--------------------")
	//字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68}
	s3 := string(slice1)
	fmt.Println(s3)

	s4 := "Welcome"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	//字符串是不能修改的
	//s4[0] = "w"
	//fmt.Println(s4)
}
