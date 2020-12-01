package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	srcFile := `D:\goTempFile\a.txt`
	file1, err := os.OpenFile(srcFile, os.O_RDWR, os.ModePerm)
	HandleErr(err)
	defer file1.Close()
	bs := make([]byte, 100, 100)
	file1.Read(bs)
	fmt.Println(string(bs))
	fmt.Println("--------------------------")

	file1.Seek(4, io.SeekStart)
	file1.Read(bs)
	fmt.Println(string(bs))
	fmt.Println("--------------------------")

	file1.Seek(3, io.SeekCurrent)
	file1.Read(bs)
	fmt.Println(string(bs))
	fmt.Println("--------------------------")

	file1.Seek(0, io.SeekEnd)
	file1.WriteString("ABC")
	fmt.Println(string(bs))
	fmt.Println("--------------------------")
}
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
