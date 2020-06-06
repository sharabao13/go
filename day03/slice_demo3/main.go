package main

import "fmt"

func main() {
	/*
		在已有的数组上创建切片
		slice := arr[start:end]
				 arr[:end],从头到end
				 arr[start:]从start到最后
		从已有的的数组上创建切片，该切片的底层数组就是当前数组
		长度是从start到end的切割的数据量
		容量是从start到数组的末尾

	*/
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 在已有的属组上创建切片
	fmt.Println("------------------------")
	s1 := a1[:5]
	s2 := a1[3:8]
	s3 := a1[:]
	s4 := a1[6:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println("==========================")
	fmt.Printf("s1长度: %d,容量：%d\n", len(s1), cap(s1)) //9
	fmt.Printf("s2长度: %d,容量：%d\n", len(s2), cap(s2)) //6
	fmt.Printf("s3长度: %d,容量：%d\n", len(s3), cap(s3)) //9
	fmt.Printf("s4长度: %d,容量：%d\n", len(s4), cap(s4)) //3

	fmt.Println("-------------修改数组中的元素-------------------")
	//修改底层数组中的数据，切片的数据随着改变
	a1[5] = 88
	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println("-------------修改切片中的元素-------------------")
	//修改切片中的元素数据，数组跟切片的数据随着改变
	s4[0] = 99
	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println("-------------切片扩容-------------------")
	//切片扩容会创建的新的内存地址，从旧的底层数组元素复制一份
	s3 = append(s3, 2, 2, 2, 2, 2)
	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
