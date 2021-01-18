package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	l1.PushBack("fist")
	l1.PushBack("22")
	fmt.Println(*l1)
}
