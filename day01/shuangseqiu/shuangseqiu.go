package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNums [6]int
	rand.Seed(time.Now().Unix())
	for j := 0; j < 5; j++ {
		for i := 0; i < 6; i++ {
			randNums[i] = rand.Intn(34)
		}
		//fmt.Println("生成的随机数:",randNums)
	}
	//fmt.Println("生成的随机数:",randNums)

	for i := 0; i < len(randNums); i++ {
		for j := 0; j < len(randNums)-1; j++ {
			if randNums[j] > randNums[j+1] {
				randNums[j], randNums[j+1] = randNums[j+1], randNums[j]
			}
		}
	}
	fmt.Println("红球号码:", randNums)
	buleNums := rand.Intn(17)
	fmt.Println("蓝球号码:", buleNums)
}
