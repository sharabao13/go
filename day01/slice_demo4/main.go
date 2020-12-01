package main

import "fmt"

func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d,cap: %d,pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	var a []int
	a = append(a[:i], append([]int{x}, a[i:]))
	a = append(a[:i], append([]int{x}, a[i:]...)...)       // 在第i个位置插入x
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第i个位置插入切片
}
