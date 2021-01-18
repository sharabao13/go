package main

import "fmt"

func main() {
	// map 声明初始化练习

	//用var来声明
	var map1 map[string]int
	//初始化数据
	map1 = map[string]int{"aa": 1, "bb": 2}
	fmt.Println(map1)

	//用make创建map
	map2 := make(map[int]int)
	// 初始化数据
	map2[1] = 1
	map2[2] = 2
	fmt.Println(map2)

	// for range 遍历map
	for k, v := range map2 {
		fmt.Println(k, v)
	}

	//删除map key delete

	delete(map1, "aa")
	fmt.Println(map1)
}
