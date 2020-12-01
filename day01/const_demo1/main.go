package main

import "fmt"

func main() {
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wenday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday)
}
