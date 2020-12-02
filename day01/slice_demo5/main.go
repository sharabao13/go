package main

import "fmt"

func main() {
	//slice1 := []int{1,100,3,4,5}
	//slice2 := []int{100,2,3}
	//fmt.Println("原切片",slice1)
	//fmt.Println("原切片",slice2)
	//copy(slice2,slice1)
	//fmt.Println("复制完切片",slice2)
	//copy(slice1,slice2)
	//fmt.Println("复制完切片",slice1)

	const elementCount = 1000
	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData
	copyData := make([]int, elementCount)
	c1 := copy(copyData, srcData)
	fmt.Println(c1)

	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0])
	//
	fmt.Println(copyData[0], copyData[elementCount-1])

	copy(copyData, srcData[4:6])
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d", copyData[i])
	}
}
