package main

import "fmt"

func main() {

	/*
		结构体比较
		结构体是值类型，如果每个字段具有可比性，则是可比较的。如果它们对应的字段相等，则认为两个结构体变量是相等的

		如果结构变量包含的字段是不可比较的，那么结构变量是不可比较的
	*/

	name1 := name{firstNmae: "Taylor", lastName: "Swift"}
	name2 := name{
		firstNmae: "Taylor",
		lastName:  "Swift",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not queal")
	}

	name2.lastName = "swift"
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not queal")
	}

	//image1 := image{data:map[int]int{
	//	0:155,
	//}}
	//image2 := image{data: map[int]int{
	//	0:155,
	//}}

	//if image1 == image2 { //包含字段不能比较
	//	fmt.Println("image1 and image2 are euqal")
	//}

}

type name struct {
	firstNmae string
	lastName  string
}

type image struct {
	data map[int]int
}
