package main

import "fmt"

func main() {
	/*
		数组array：存储一组相同数据类型的数据
			特点： 定长

		切片slice: 通数组一致，也叫变长数组或动态数组
			特点 ：变长
			是一个引用类型的的容器，指向了一个底层数组

		切片定义
			var [] int {}  //[] 表示切片  [4/...]表示数组
	*/

	// 定义一个数组
	arr1 := [3]int{1, 2, 3} //[3] 有长度的表示数组
	fmt.Printf("数组的值是: %d\n", arr1)
	fmt.Printf("数组的类型是: %T\n", arr1)

	// 定义一个切片
	var s1 []int
	fmt.Printf("切片的值是: %d\n", s1)
	fmt.Printf("切片的类型是: %T\n", s1)

	/*
		内建函数 make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，跟 new 类似，
		第一个参数也是一个类型而不是一个值，跟 new 不同的是，make 返回类型的结构实列而不是指针，而返回值也依赖于具体传入的类型
		func make(t Type,size IntegerType) Type
			ex:make([]int,0,10) //0 长度  //10 容量
		// 超过切片的长度会自动扩容

	*/

	// make 创建一个切片

	S3 := make([]int, 2, 5)
	fmt.Println(S3)
	fmt.Printf("切片S3的长度: %d,容量是: %d\n", len(S3), cap(S3))
	// 切片添加值
	S3[0] = 1
	S3[1] = 2
	fmt.Println(S3)

	/*
			切片追加成员值
			append(type,len,cap)
		该函数第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间
	*/
	//append 追加成员值
	s4 := make([]int, 0, 5)
	s4 = append(s4, 1, 2, 3)
	fmt.Println(s4)
	fmt.Println(len(s4))
	s4 = append(s4, 4, 5, 6, 7, 8)
	fmt.Println(s4)
	fmt.Println(len(s4))
	// s4的基础上最近S3成员值 用...表示
	s4 = append(s4, S3...)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println("---------------")

	//切片的遍历通数组一致

	for i := 0; i < len(s4); i++ {
		fmt.Print(s4[i], "\t")
	}
	fmt.Println()
	fmt.Println("---------------")

	// for range 变量切片
	for _, val := range s4 {
		fmt.Print(val, "\t")
	}
	fmt.Println()
}
