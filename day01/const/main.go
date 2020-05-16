//常量 const

package main

import "fmt"

// 全局常量声明
const pi = 3.1592654
const e = 2.7182

// 批量常量声明
//const(
//	c1 = 1
//	c2 = 2
//)
func main() {
	//批量声明常量，如常量值省略，值同上一行一致
	//const (
	//	c1 = 100
	//	c2
	//	c3 = 101
	//	c4
	//)

	// iota 是go语言的常量计数器，只能在常量的表达式中使用,
	// iota 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
	// iota 出现 _ 时 跳过该累计值
	//const (
	//	c1 = iota //c1 = 0
	//	c2        // 1
	//	_         // 2
	//	c3        // 3
	//	c4        // 4
	//)

	// 声明常量中出现多个iota
	//const (
	//	c1 = iota    //0
	//	c2 = 3.15    //3.15  //iota 1
	//	c3 = iota    //2
	//	c4           //3
	//)

	// 定义数量级别
	//const (
	//	_ = iota
	//	KB = 1 << (10 * iota)
	//	MB = 1 << (10 * iota)
	//	GB = 1 << (10 * iota)
	//	TB = 1 << (10 * iota)
	//	PB = 1 << (10 * iota)
	//)

	// 多个iota定义在一行

	const (
		c1, c2   = iota + 1, iota + 2 // 1,2
		c3, c4                        // 2,3
		c5, c6   = 6, 7               // 6,7
		c7, c8                        // 6,7
		c9, c10  = iota + 1, iota + 2 // 5,6
		c11, c12                      // 6,7
	)

	//fmt.Println(KB)
	//fmt.Println(MB)
	//fmt.Println(GB)
	//fmt.Println(TB)
	//fmt.Println(PB)
	fmt.Println(pi)
	fmt.Println(e)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("c5:", c5)
	fmt.Println("c6:", c6)
	fmt.Println("c7:", c7)
	fmt.Println("c8:", c8)
	fmt.Println("c9:", c9)
	fmt.Println("c10:", c10)
	fmt.Println("c11:", c11)
	fmt.Println("c12:", c12)

}
