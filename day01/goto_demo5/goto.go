package main

import "fmt"

//goto 用法
func main() {
	sum := 0
	i := 1
START:
	if i > 100 {
		goto END
	}
	sum += i
	i++
	goto START
END:
	fmt.Println(sum)
}
