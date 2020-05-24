package main

import "fmt"

func main() {
	/*
		if 嵌套练习
	*/
	score := 0
	fmt.Println("请输入一个成绩: ")
	fmt.Scanf("%d", &score)
	if score < 60 {
		fmt.Printf("您输入的成绩是: %d,成绩不及格\n", score)
	} else {
		if score > 100 {
			fmt.Println("请输入100以内的成绩分数")
		} else {
			if score >= 80 {
				fmt.Printf("您输入的成绩是: %d,成绩良好\n", score)
			} else {
				fmt.Printf("您输入的成绩是: %d,成绩及格\n", score)
			}
		}
	}
}
