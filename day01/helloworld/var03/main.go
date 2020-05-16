// 变量声明练习
package main

import "fmt"

var (
	id     int
	name   string
	gender bool
)

func main() {
	id = 1
	name = "bao"
	gender = true
	fmt.Println(id)
	fmt.Printf("name:%s\n", name)
	fmt.Println(gender)
}
