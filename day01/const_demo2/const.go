package main

import "fmt"

func main() {
	const name string = "Tae"
	fmt.Println(name)
	// name = "Yeon"
	fmt.Println(name)

	const (
		NAME   = "Yeon"
		AGE    = 18
		GENDER = "MALE"
	)
	fmt.Println(NAME, AGE, GENDER)

	const PI = 3.1315
	fmt.Println(PI)

	const A, B, C, D int = 1, 2, 3, 4
	fmt.Println(A, B, C, D)

	const (
		A1 = iota
		B1
		C1
		D1
		E1 = iota
		F1
	)
	fmt.Println(A1, B1, C1, D1, E1, F1)
	fmt.Println("-----------------------")
	const (
		_ = iota
		X1
		X2
		X3
		X4
		X5 = 100
		X6
	)
	fmt.Println(X1, X2, X3, X4, X6)
}
