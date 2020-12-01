package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		os包提供了操作系统函数的不依赖平台的接口。设计为Unix风格的，虽然错误处理是go风格的；失败的调用会返回错误值而非错误码。
	通常错误值里包含更多信息。例如，如果某个使用一个文件名的调用（如Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为*PathError，
	其内部可以解包获得更多信息
	 */

	//fileInfo
	srcFile := `D:\goTempFile\a.png`
	file,err := os.OpenFile(srcFile,os.O_RDWR,os.ModePerm)
	HandleErr(err)
	fileInfo,err := os.Stat("D:\\goTempFile\\a.png")
	HandleErr(err)
	fmt.Println(file.Name())
	fmt.Println(fileInfo.)
	fmt.Println(fileInfo.Size())
}
func HandleErr(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}
