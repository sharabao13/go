// 浮点数
// Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
// float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.1f\n", math.Pi) //%.1f 数值代表打印几位小数
	a := 10.343331
	b := 3.3123
	fmt.Printf("%.4f\n", a+b) // 13.6556
	fmt.Printf("%.2f\n", a+b) // 13.66
}
