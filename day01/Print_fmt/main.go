// fmt.Printf 格式化输出
// 占位符
// %T 查看数据类型
// %v 查看变量值
// %#v 打印的值带类型
// %b 二进制
// %d 十进制
// %o 八进制
// %x 十六进制
// %f 浮点数
// %.2f 打印多少位小数点
// %s 字符串

package main

import "fmt"

func main() {
	i := 2.23676
	s := "test"
	fmt.Printf("i的数据类型：%T\n", i)
	fmt.Printf("i的值位：%v\n", i)
	fmt.Printf("i的值位：%#v\n", i)
	fmt.Printf("i的二进制值位：%b\n", i)
	fmt.Printf("i的八进制值位：%o\n", i)
	fmt.Printf("i的十进制值位：%d\n", i)
	fmt.Printf("i的十六进制值位：%x\n", i)
	fmt.Printf("i的浮点数值位：%f\n", i)
	fmt.Printf("i的浮点数2为值位：%.2f\n", i)
	fmt.Printf("s的数据类型：%T\n", s)
	fmt.Printf("s的值位：%#v\n", s)
	fmt.Printf("s的值位：%s\n", s)
}
