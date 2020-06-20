package main

import "fmt"

func main() {
	/*
		1.函数返回值  一个函数被调用后，返回给调用处的执行结果
			调用处需使用变量接收该结果
			一个函数可以返回多个值
		2.return 语句
			在定义一个函数有返回值是，必须使用return语句，将结果返回调用处

	*/
	//有返回值的函数需赋值给变量
	res := getSum()
	fmt.Println("总和是：", res)
	fmt.Println("------------------")
	res1 := getSum1()
	fmt.Println("总和是：", res1)
}

//返回值用法1
func getSum() int { //int调用处的结
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum //
}

//函数返回值用法2
func getSum1() (sum int) { //(sum,int) 调用处语法，返回值名称，跟数据类型
	sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return // 没写返回值，默认调用处定义的名称
}
