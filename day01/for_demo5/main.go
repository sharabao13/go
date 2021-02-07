package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 乘法表
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
	/*猜数字游戏
	  	生成一个[0-100]随机数
	      让用户最多猜5次（从命令行输入一个整数）
	  	猜太大 提示猜太大还有几次机会
	      猜调小提示猜太小还有几次机会
	      猜中 提示猜中了
	  	5次多没中退出
	*/
	//随机函数种子
	rand.Seed(time.Now().Unix())
	//定义一个0-100随机数
	ranDom1 := rand.Intn(101)
	fmt.Println(ranDom1)
	//定义一个接收用户输入数字的变量
	var num1 int
	//定义次数
	sum := 5
START:
	fmt.Print("猜大小游戏开始，请输入一个整数: ")
	fmt.Scan(&num1)
	for i := 1; i <= 5; i++ {
		if num1 == ranDom1 {
			fmt.Println("太厉害了，猜中了")
			break
		}
		if num1 > ranDom1 {
			sum -= i
			fmt.Printf("猜错了！输入的数字大于随机数,还有%d次机会\n", sum)
			if sum == 0 {
				fmt.Println("太笨了5次机会没猜中")
				break
			}
			goto START
		} else if num1 < ranDom1 {
			sum -= i
			fmt.Printf("猜错了！输入的数字小于于随机数,还有%d次机会\n", sum)
			if sum == 0 {
				fmt.Println("太笨了5次机会都没猜中")
				break
			}
			goto START
		}

	}
}
