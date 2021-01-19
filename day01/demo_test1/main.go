package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Input You Name：")

	//读取数据知道\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error accourrd: %s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("Hello %s! What can i do fo you?\n", name)
	}

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error accourrd: %s\n", err)
			//异常退出
			continue
		}
		input = input[:len(input)-2]

		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "hello":
			fmt.Println("Hello")
			os.Exit(0)
		case "nothing", "bye":
			fmt.Println("bye")
			os.Exit(1)
		default:
			fmt.Println("Sorry,i did't Catch You")
		}

	}
}
