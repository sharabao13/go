package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map 遍历
		for range

	*/
	map1 := make(map[int]string)
	map1[0] = "apple"
	map1[1] = "banana"
	map1[2] = "pear"
	map1[3] = "txxxx"
	map1[4] = "555555"
	map1[5] = "fdsiadf"
	map1[6] = "fepkadkfdsfsa"
	for k1, v1 := range map1 {
		fmt.Println(k1, v1)
	}

	//按顺序打印map元素
	for i := 0; i < len(map1); i++ {
		fmt.Println(i, map1[i])
	}

	// map是无序的 按升序/降序打印map
	/*
		1.获取所有key的值，素组或切片
		2.进行排序
		3.遍历key
	*/

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	for k1, _ := range map1 {
		keys = append(keys, k1)
	}
	fmt.Println(keys)
	//for ks, vs := range keys {
	//	fmt.Println(ks, vs)
	//}
	//fmt.Println(len(keys))
	//冒泡排序的方法按大小排序keys值
	//for i := 0; i < len(keys); i++ {
	//	for j := 0; j < len(keys)-1; j++ {
	//		if keys[j] > keys[j+1] {
	//			keys[j], keys[j+1] = keys[j+1], keys[j]
	//		}
	//	}
	//}
	//fmt.Println(keys)

	// sort 方法
	// sort.Ints(keys) 整形排序
	// sort.String(keys) 字符串排序
	//s1 := []string{"abc", "ADD", "abed", "看到"}
	//fmt.Println(s1)
	//sort.Strings(s1)
	//fmt.Println(s1)

	sort.Ints(keys)
	fmt.Println(keys)

	for _, k2 := range keys {
		fmt.Println(k2, map1[k2])
	}

}
