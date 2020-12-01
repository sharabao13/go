package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sara", "August", "Heidi", "Emilie", "Peter", "Glana", "Adriano", "Elizabeth"}
	length       = len(users)
	distribution = make(map[string]int, length)
)

func dispatchCoin() (left int) {
	var c int                 // 设置一个变量存储每次分配的金币数
	for _, v := range users { // 1. 获取人名
		c = 0
		for _, vv := range v { // 2. 按规则分配，遍历人名中的每个字符
			vm := string(vv)
			if vm == "e" || vm == "E" {
				c += 1
			}

			if vm == "i" || vm == "I" {
				c += 2
			}

			if vm == "o" || vm == "O" {
				c += 3
			}

			if vm == "u" || vm == "U" {
				c += 4
			}
		}
		distribution[v] = c // 每个人分的金币数保存到 distribution 中
		coins -= c          // 记录每次分配后剩下金币数
	}
	left = coins // 记录最终剩余金币数
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println(distribution)
	fmt.Println("剩下：", left)
}
