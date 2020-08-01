package main

import (
	"fmt"
	"net"
)

func main() {

	/*
		错误类型处理
	*/
	addr, err := net.LookupHost("www.jjjsdddddddss.com")
	fmt.Println(err)
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("DNS超时")
		} else if ins.IsTemporary {
			fmt.Println("临时性错误")
		} else if ins.IsNotFound {
			fmt.Println("未找到")
		} else {
			fmt.Println("其他错误")
		}
	}
	fmt.Println(addr)
}
