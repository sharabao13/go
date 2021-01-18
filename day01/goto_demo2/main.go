package main

import "fmt"

func main() {
	var breakAgain bool
	for y := 0; y <= 10; y++ {
		for x := 0; x <= 10; x++ {
			if x == 5 {
				breakAgain = true
				break
			}
		}
		if breakAgain {
			break
		}
	}
	fmt.Println("done")

}
