package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片
	//1. var 定义切片
	var s1 []int
	fmt.Println(s1)

	//make 函数
	s2 := make([]int, 3)
	fmt.Printf("%v,%d,%d\n", s2, len(s2), cap(s2))
	// append 函数 在元素后面增加数据
	s2 = append(s2, 0, 1, 2, 3, 4)
	fmt.Printf("%v,%d,%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 5, 9)
	fmt.Printf("%v,%d,%d\n", s3, len(s3), cap(s3))

	s4 := make([]int, 3)

	//元素 增 删 改 查
	//增
	s4 = append(s4, 0, 1, 2)
	fmt.Println(s4)

	//删除
	// copy 函数
	s6 := []int{1, 2, 3}
	s7 := []int{10, 20, 30, 40}
	//copy(s7,s6)
	//fmt.Printf("copy完s7的数据%v\n",s7)
	ss1 := []string{"a", "b", "c"}
	ss2 := []string{"test", "go", "function", "software"}

	fmt.Println("切片删除")
	//切片删除
	s7 = s7[:2]
	fmt.Println(s7)

	s8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(s8[3:], s8[4:])
	fmt.Println(s8[:len(s8)-1])

	fmt.Println("--------------------")

	copy(s6, s7)
	fmt.Println(s6)

	fmt.Println("字符串copy")
	copy(ss2, ss1)
	fmt.Println(ss2)
	//改
	s4[0] = 100
	s4[1] = 90
	s4[2] = 80
	fmt.Println(s4)

	//查
	fmt.Println(s4[3])

	var s5 []int
	s5 = s4
	fmt.Printf("修改前s5的值: %v\n", s5)
	fmt.Println("---------------------")
	s5[3] = 70
	fmt.Printf("修改后s5的值: %v\n", s5)
	fmt.Println("---------------------")
	fmt.Printf("修改前s4的值: %v\n", s4)

	//队列 每次添加在队尾，移除元素在队头
	fmt.Println("队列")
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	fmt.Println(queue)
	fmt.Println(queue[1:])
	fmt.Println(queue[2:])
	fmt.Println(queue[3:])

	// 堆栈 每次添加在队尾，移除元素在队尾
	fmt.Println("堆栈")
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack)
	fmt.Println(stack[:len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[:len(stack)-1])

	//多维切片

	point := [][]int{}
	point = append(point, []int{1, 2, 3})
	point = append(point, []int{5, 6, 7})
	point = append(point, []int{0, 0, 0})
	point = append(point, []int{0, 0, 0})
	fmt.Println(point)
	fmt.Println(point[3][1])

	//sort 切片排序
	//整形类型的排序 ints  从小到大
	ss3 := []int{8, 1, 100, 22, 13, 9}
	sort.Ints(ss3)
	fmt.Println(ss3)

	//字符串排序
	ss4 := []string{"test", "bao", "strut", "fuction"}
	sort.Strings(ss4)
	fmt.Println(ss4)

}
