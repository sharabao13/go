package main

import "fmt"

func main() {
	var arr1 [2][3][4]int
	fmt.Println(arr1)

	arr1[0][1][3] = 4
	fmt.Println(arr1)
	fmt.Println("--------------------------------")

	var arr2 [2][3]int
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("请输入第%d行第%d列1的数\n", i+1, j+1)
			fmt.Scanln(&arr2[i][j])
		}
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "\t")
		}
		fmt.Print("\n")
	}

	//for _,v := range arr2{
	//	fmt.Println(v)
	//}
}
