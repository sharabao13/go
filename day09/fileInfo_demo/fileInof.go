package main

import (
	"fmt"
	"os"
)

func main() {
	/* fileInfo 文件信息
	os.stat 包下的方法
	*/
	fileInfo, err := os.Stat(`E:\\untitled\\go_ducment\a.txt`)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	fmt.Println("文件名是:", fileInfo.Name())
	fmt.Println("文件大小: ", fileInfo.Size())
	fmt.Println("是否目录: ", fileInfo.IsDir())
	fmt.Println("文件修改时间", fileInfo.ModTime())
}
