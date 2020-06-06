package main

import "fmt"

func main() {
	/*
		map[key]string
			string 也可以是map
		map[string]map[string]tring
	*/
	map1 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "li"
	m1["age"] = "18"
	map1["hr"] = m1
	fmt.Println(map1)

	Meg2 := make(map[string]string)
	Meg2["name"] = "tae"
	Meg2["age"] = "18"
	map1["leader"] = Meg2
	fmt.Println(map1)
}
