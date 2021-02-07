package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机数
func main() {
	num1 := rand.Int()
	fmt.Println(num1)

	num2 := rand.Intn(10)
	fmt.Println(num2)

	rand.Seed(time.Now().Unix())
	num3 := rand.Intn(101)
	fmt.Println(num3)
}
