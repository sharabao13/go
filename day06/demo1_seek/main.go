package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 断点续传
func main() {
	// 定义文件目录
	srcFile := `D:\goTempFile\a.png`
	destFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]
	//fmt.Println(destFile)
	tempFile := destFile + "tempFile.txt"
	//fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	HandleErr(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	//step1 读取临时文件中的数据，在seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	//设置读 写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1034)
	n2 := -1
	n3 := -1
	total := int(count)

	//复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕", total)
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		//复制的总量存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Printf("复制的总量total: %d\n", total)
		//
		//if total > 2000{
		//	panic("系统异常")
		//}
	}

}
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
