package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生产随机数
	var randNums [10]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		randNums[i] = rand.Intn(101)
	}
	fmt.Println("生成的随机数:", randNums)
	//冒泡排序
	for i := 0; i < len(randNums); i++ {
		for j := 0; j < len(randNums)-1; j++ {
			if randNums[j] > randNums[j+1] {
				randNums[j], randNums[j+1] = randNums[j+1], randNums[j]
			}
		}
	}
	fmt.Println("排序完的随机数:", randNums)
}
