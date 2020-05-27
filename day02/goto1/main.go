package main

import "fmt"

func main() {
	/*
		goto 语句

	*/
	for a := 1; a < 10; a++ {
	LOOP:
		if a == 6 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a的值: %d\n", a)
	}
}
