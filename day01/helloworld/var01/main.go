package main

import "fmt"

var (
	studentName string
	age         int
	gender      string
	id          int
	isOk        bool
)

func main() {
	id = 1
	studentName = "Li"
	age = 18
	gender = "Man"
	isOk = false
	fmt.Print(id)
	fmt.Println()
	fmt.Print(age)
	fmt.Println()
	fmt.Printf("studentName:%s\n", studentName)
	fmt.Println(gender)
	fmt.Print(isOk)
}
