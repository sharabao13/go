package main

import (
	"fmt"
	"sort"
)

func main() {
	sence := make(map[string]float32)
	sence["one"] = 1.22
	sence["two"] = 3.12
	sence["three"] = 4.12

	for k, v := range sence {
		fmt.Println(k, v)
	}
	fmt.Println("-----------------------")

	scene := make(map[string]int)
	scene["zero"] = 0
	scene["dest"] = 1
	scene["banana"] = 3
	scene["apple"] = 5
	var sceneList []string
	for k, _ := range scene {
		sceneList = append(sceneList, k)
	}
	fmt.Println(sceneList)
	sort.Strings(sceneList)
	fmt.Println(sceneList)

	//删除map key value

	var testMap1 map[int]int
	testMap1 = map[int]int{1: 1, 2: 2}
	fmt.Println(testMap1)
	delete(testMap1, 1)
	fmt.Println(testMap1)

}
