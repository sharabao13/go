package main

import (
	"log"
	"os"
)

func main() {
	destFile := "test.txt"
	file, err := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	LogErr(err)
	defer file.Close()
	file.WriteString("kk")

}
func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
