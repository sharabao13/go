package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)

	a2 := [...]string{"s1", "s2", "s3", "s4", "s5"}
	for _, v1 := range a2 {
		fmt.Println(v1)
	}
	fmt.Println("------------------------------")
	//3 := []
}
