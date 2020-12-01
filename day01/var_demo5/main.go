package main

import "fmt"

func demoTest(a int) int {
	var b int
	b = a
	return b
}

func void() {

}

func main() {
	var a int = 10
	void()
	fmt.Println(a, demoTest(a))
}
