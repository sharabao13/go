package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := `D:\src_dir\ebsgit_all.sh`
	destFile := `D:\dest_dir\a.sh`
	c1, err := copyFile(srcFile, destFile)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(c1)

}
func copyFile(srcFile, destFile string) (int, error) {
	sFile, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	dFile, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	defer sFile.Close()
	defer dFile.Close()

	ns := make([]byte, 1024, 1024)
	Total := 0
	for {
		n, err := sFile.Read(ns)
		if err == io.EOF {
			fmt.Println("文件复制完毕")
			break
		} else if err != nil {
			fmt.Println("err", err)
		}
		Total += n
		dFile.Write(ns[:n])
	}
	return Total, nil

}
