package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		文件写入
		writer Write(b []byte) (n int, err error)
	*/
	//写入本地文件中fileName2 := `E:\\untitled\\go_ducment\b.txt`
	fileName2 := `E:\\untitled\\go_ducment\b.txt`
	//step1 打开文件 如文件不存在创建文件
	file, err := os.OpenFile(fileName2, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	//step3 关闭文件
	defer file.Close()

	//写入文件
	bs := []byte{97, 98, 99, 100}
	n, err := file.Write(bs)
	fmt.Println(err)
	fmt.Println(n)
	HandeErr(err)
	file.WriteString("\n")

	//直接写出字符串wirtestring WriteString(s string) (n int, err error)
	n, err = file.WriteString("Hello World")
	fmt.Println(n)
	HandeErr(err)
	file.WriteString("\n")

}
func HandeErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
