package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.4555
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok {
			fmt.Printf("半径是: %.2f", err.radius)
		}
		return
	}
	fmt.Println("圆形的面积是: ", area)
}

// 1.定义一个结构体表示错信息
type areaError struct {
	msg    string
	radius float64
}

//实现error接口方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error: 半径：%.2f,%s", e.radius, e.msg)
}

//定义求圆半径的函数
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法的", radius}
	}
	return math.Pi * math.Pow(radius, 2), nil
}
