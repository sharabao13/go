package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	fmt.Println("生成的原数组为:", arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println("排序完成后的数组为:", arr)
}
