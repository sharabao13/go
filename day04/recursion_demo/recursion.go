package main

import "fmt"

func add(n int) int {
	if n == 1 {
		return 1
	}
	return n + add(n-1)
}
func multi(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * multi(n-1)
}

func tower(A, B, C string, n int) {
	if n == 1 {
		fmt.Printf("%s--->%s\n", A, C)
		return
	} else {
		tower(A, C, B, n-1)
		fmt.Printf("%s--->%s\n", A, C)
		tower(B, A, C, n-1)
	}

}
func main() {
	fmt.Println(add(100))
	fmt.Println(multi(0))
	tower("A", "B", "C", 5)
}
