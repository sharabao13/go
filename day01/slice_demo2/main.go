package main

import "fmt"

func main() {
	var numbers []int
	fmt.Println(numbers)

	numbers = append(numbers, 0)
	fmt.Println(numbers)
	if numbers == nil {
		fmt.Println("空切片")
	}
	numbers = append(numbers, 1)
	fmt.Println(numbers)

	numbers = append(numbers, 2, 3, 4)
	fmt.Println(numbers)

	s1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var s2 []int
	for _, k := range s1 {
		if k == 5 {
			continue
		}
		fmt.Println(k)
	}
	s2 = append(s1[0:4])
	fmt.Println(s2)

}
