package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Stat("2.go")
	file2, _ := os.Stat("2.go")
	if os.SameFile(file1, file2) {
		fmt.Println("文件一样")
		return
	}
	fmt.Println("文件不一样")
}
