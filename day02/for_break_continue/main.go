package main

import "fmt"

func main() {
	/*
				循环介绍
					1.循环条件不满足时，终止循环
					2.可以通过break，continue强制结束循环

					break ：彻底结束循环，终止
			     	continue: 结束某个循环，下次继续

		  注意点多重嵌套循环，break，continue默认结束的里层循环
			如果想指定的结束某个循环，可以给循环加标签（起名）
	*/
	//for i := 1; i < 10; i++ {
	//	if i == 5 {
	//		//break //强制结束循环
	//		continue // 跳过i == 5继续执行循环体
	//	}
	//	fmt.Println(i)
	//}

	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	// 都重循环continue,break默认里层循环结束

out:
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			if b == 2 {
				continue out
				//break out
			}
			fmt.Printf("%d,%d\n", a, b)
		}
	}
}
