package main

import "fmt"

func main() {
	/*
		map 映射，专门用于存储键值对的集合，属于引用类型

		存储特点:
			1.存储是无序的
			2.键不能重复，并且和值一一对应
				map中的键重复不会报错，但值会覆盖原来键的值

		语法结构：
			1.创建map
				var map1 map[key类型]value类型 //声明的map如不赋值将不能使用
					nil map 无法直接使用

				var map2 = make(map[key类型]value类型) 可以直接使用

				var map3 = map[key类型]value类型{key:value,key:value...}

	*/

	// 创建map

	var map1 map[int]string //niu=l 没有初始化

	fmt.Println(map1)
	fmt.Println(map1 == nil)

	var map2 = make(map[int]string)
	fmt.Println(map2)
	fmt.Println(map2 == nil)

	var map3 = map[string]int{"chinese": 60, "math": 80, "english": 100}
	fmt.Println(map3)

	var map4 map[string]int
	if map4 == nil {
		map4 = make(map[string]int)
		fmt.Println(map4 == nil)
	}

	// 添加map元素
	// map[key] = value
	map2[1] = "acd"
	map2[2] = "txxx"
	map2[3] = "tae"
	map2[4] = "bac"
	map2[6] = ""
	fmt.Println(map2)
	//获取map元素
	fmt.Println(map2[1])
	/* value,ok=map[key]
	ok = true 值存在
	ok = false 值不存在
	*/
	v1, ok := map2[1]
	if ok {
		fmt.Printf("的值是: %s\n", v1)
	} else {
		fmt.Printf("map2[7]的是值是空值: %s\n", v1)
	}

	//修改map的值
	map2[0] = "000000"
	map2[1] = "string"
	fmt.Println(map2)

	//删除map 元素 delete
	//delete(map,key)

	delete(map2, 3)
	fmt.Println(map2)

	// map 长度 len方法
	fmt.Println(len(map2))

	scort := map[string]int{"语文": 90, "数学": 89}
	for k, v := range scort {
		fmt.Printf("%s的分数是: %d\n", k, v)
	}

}
