package main

import "fmt"

func main() {

	/*
			深拷贝：拷贝数据本身
				值类型的数据都是深拷贝: array,int float bool string struct
		    浅拷贝：拷贝的数据地址
				引用类型的数据都是浅拷贝： slice map

	*/

	s1 := []int{1, 2, 3}
	s2 := make([]int, 0)
	for _, val := range s1 {
		s2 = append(s2, val)
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 0
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("--------------")
	//copy 切片深拷贝
	//func copy(dst, src []Type) int
	//copy(目标切片，源切片)
	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("----------------")
	copy(s3, s2)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("--------------")

	s4 := []int{5, 6, 7, 8}
	//从目标切片拷贝数据到源切片，目标切片的容量如果大于源切片只拷贝源切片容量的大小的数据
	copy(s3, s4[:])
	fmt.Println(s3) //1278
	fmt.Println(s4)
}
