package main

import "fmt"

func main() {
	/*
		函数多返回值
	*/

	rentanglePer, rentangleArea := rentangle(2.86, 4.858)
	fmt.Printf("长方形的周长是：%.2f,面积是： %.2f\n", rentanglePer, rentangleArea)

}

// 创建一个计算长方形周长/面积的函数
func rentangle(a, b float64) (perimeter float64, area float64) {
	perimeter = (a + b) * 2
	area = a * b
	return
}
