package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		复制文件
		step1 读取源文件
		step2从源文件遍历切片数据到目标文件
		step3 关闭文件
	*/
	srcFile := `E:\\untitled\\go_ducment\1200px-Go_Logo_Blue.svg.png`
	destFile := `E:\\untitled\\go_ducment\b1.png`
	//total,err := copyFile(srcFile,destFile)
	//if err != nil {
	//	fmt.Println("err",err)
	//}
	//fmt.Println(total)

	w1, err := copyFile2(srcFile, destFile)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(w1)

}

//io.copy 复制文件用法
func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
		return 0, nil
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

// 传统复制文件方法
func copyFile(srcFile, destFile string) (int, error) { //源文件和目标文件参数，返回读取的数据量跟错误信息
	//step1 打开源文件 和目标文件
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm) //目标文件如果不存在创建文件
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	//step4 关闭文件
	defer file1.Close()
	defer file2.Close()

	//读写复制文件
	bs := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := file1.Read(bs)
		if err == io.EOF {
			fmt.Println("复制完毕")
			break
		} else if err != nil {
			fmt.Println("err", err)
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}
