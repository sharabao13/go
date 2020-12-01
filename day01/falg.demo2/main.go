package main

import (
	"flag"
	"fmt"
)

var testDataString = flag.String("name", "Tae", "Input Your Name")
var testDataInt = flag.Int("age", 18, "Input Your Age")

func main() {
	flag.Parse()
	fmt.Println(*testDataString)
	fmt.Println(*testDataInt)

	str := new(string)
	*str = "Taetae@gamil.com"
	fmt.Println(*str)

	a, b := 1, 3
	value := calc(a, b)
	fmt.Println(value)
}

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x

}
