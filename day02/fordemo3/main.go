package main

import "fmt"

func main() {
	// 打印58-23的数值
	for Num := 58; Num >= 23; Num-- {
		fmt.Println(Num)
	}

	// 打印1-100的和
	res := 0
	for i := 1; i <= 100; i++ {
		res = res + i
	}
	fmt.Printf("1-100的和是: %d\n", res)

	//打印1-100内，能过被3整除，但是不能被5整除的数值，统计别打印数字的个数，每行打印5个

	res1 := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			res1++
			fmt.Print("\t", i)
			if res1%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Printf("能被3整除不能被5整除的个数是: %d\n", res1)
}
