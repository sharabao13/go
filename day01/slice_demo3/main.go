package main

import (
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	fmt.Println("-----------------------")

	var highRishBuilding [30]int
	for i := 0; i < 30; i++ {
		highRishBuilding[i] = i + 1
	}
	//区间
	fmt.Println(highRishBuilding[10:15])

	//中间到尾部
	fmt.Println(highRishBuilding[20:])

	//从开头指定位置
	fmt.Println(highRishBuilding[:11])
	fmt.Println(highRishBuilding)
	//重置
	fmt.Println(highRishBuilding[0:0])

	fmt.Println("-------------------------")

	//声明字符串切片
	var strList []string

	//声明整数切片
	var intList []int

	var numListEmpty = []int{}

	//输出三个切片
	fmt.Println(strList, intList, numListEmpty)

	var numList1Empty = []int{}
	if numList1Empty == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	//make 创建切片
	c := make([]int, 2)
	d := make([]int, 2, 10)
	fmt.Println(c, d)

}
