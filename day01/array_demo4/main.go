package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 2, 3, 4, 5}

	for k, v := range a1 {
		fmt.Println(k, v)
	}
}
