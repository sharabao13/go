package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -16.6
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("area: %.2f", area)
}

//传教计算圆的面积函数

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("radius less than zero: %.2f", radius)
	}
	return math.Pi * math.Pow(radius, 2), nil
}
