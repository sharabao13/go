package main

import "fmt"

func main() {
	a, b := 1, 2
	swapPtr(&a, &b)
	fmt.Println(a, b)
}

func swapPtr(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
