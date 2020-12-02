package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigend map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigend = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415
	mapAssigend["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}
	fmt.Println(noteFrequency)
	map1 := make(map[int][]int)
	map2 := make(map[int]*[]int)
	fmt.Println(map1, map2)

}
