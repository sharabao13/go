package main

import (
	"fmt"
	"math"
)

//定义一个map函数
func fun3(m1 map[string]int) {
	fmt.Println("函数中修改前map的数据", m1)
	m1["one"] = 100
	fmt.Println("函数中修改后map的数据", m1)

}

//定义一个切片函数
func fun2(s1 []int) {
	fmt.Println("函数中切片的数据", s1)
	s1[0] = 100
	fmt.Println("函数中修改后切片的数据", s1)

}

//定义数组函数
func fun1(a1 [4]int) {
	fmt.Println("函数中数组的数据: ", a1)
	a1[0] = 10
	fmt.Println("函数中修改后数组的数据: ", a1)
}

func Oddeven(num int) (odd, even string) {
	if num%2 == 1 {
		odd = "奇数"
		return
	} else {
		even = "偶数"
		return
	}
}

func ret(x, y float64) (perimeter, area float64) {
	perimeter = (x + y) * 2
	area = x * y
	return
}

func CircleArea(r float64) float64 {
	area := math.Pi * math.Pow(r, 2)
	return area
}

func main() {
	Area := CircleArea(3)
	fmt.Printf("圆的面积是: %.2f\n", Area)

	per, area := ret(3, 4)
	fmt.Printf("矩形的周长是: %.2f\n", per)
	fmt.Printf("矩形的周长是: %.2f\n", area)

	odd, even := Oddeven(3)
	fmt.Println(even)
	fmt.Println(odd)
	a1 := [4]int{1, 2, 3, 4}
	fmt.Println("调用前数组数据", a1)
	fun1(a1)
	fmt.Println("调用后数组数据", a1)
	fmt.Println("--------------------")
	s2 := []int{1, 2, 3, 4}
	fmt.Println("调用前切片的数据", s2)
	fun2(s2)
	fmt.Println("调用后切片数据", s2)
	fmt.Println("---------------------")
	var m1 map[string]int
	m1 = map[string]int{"one": 1, "two": 2}
	fmt.Println("调用前map的数据", m1)
	fun3(m1)
	fmt.Println("调用后map数据", m1)

}
