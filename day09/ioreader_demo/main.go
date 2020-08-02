package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		文件读操作
		read接口
		将文件的数据读取为一个字节片。读取和写入从参数切片的长度中获取字节数
		Read(b []byte) (n int, err error)
	*/

	//读取本地文件中fileName2 := `E:\\untitled\\go_ducment\a.txt`
	// step1 打开文件 建立程序与文件的连接
	fileName := `E:\\untitled\\go_ducment\a.txt`
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	//step3 关闭文件
	defer file.Close()

	//step2 读取文件
	bs := make([]byte, 4, 4)
	//第一次读取数据
	/*
		n,err := file.Read(bs)
		if err != nil {
			fmt.Println("err",err)
			return
		}
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		//第二次读取数据
		n,err = file.Read(bs)
		if err != nil {
			fmt.Println("err",err)
			return
		}
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		//第三次读取数据
		n,err = file.Read(bs)
		if err != nil {
			fmt.Println("err",err)
			return
		}
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

		//第四次读取数据
		n,err = file.Read(bs)
		if err != nil {
			fmt.Println("err",err)
			return
		}
		fmt.Println(n)
		fmt.Println(bs)
		fmt.Println(string(bs))

	*/
	// 遍历文件
	//n := -1
	for {
		n, err := file.Read(bs)
		if err == io.EOF {
			fmt.Println("读到文件末尾,EOF")
			break
		}
		fmt.Println(string(bs[:n]))
	}

}
