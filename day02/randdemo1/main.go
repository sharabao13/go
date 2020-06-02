package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		随机数练习
	*/
	//for i := 1; i <= 5; i++ {
	//	fmt.Println(rand.Intn(10))
	//}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		fmt.Println(rand.Intn(5))
	}
}
