package main

import "fmt"

func main() {
	str1 := "abcdef"
	ret1 := simpleHase(str1)
	fmt.Println(ret1)

}

func simpleHase(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return
}
