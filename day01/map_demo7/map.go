package main

import "fmt"

func main() {
	/*
		map key value 存储  无序的
	*/

	// 1. 定义
	var m1 map[int]string
	fmt.Printf("%T,%v\n", m1, m1) //类型和值

	//make 函数定义
	m2 := make(map[string]int)
	fmt.Printf("%T,%v\n", m2, m2) //类型和值
	//fmt.Println(m2 == nil) //0值 不等空

	// 初始化
	m1 = map[int]string{1: "english", 2: "chinese"}
	fmt.Println(m1)

	m2["tae"] = 1
	m2["bao"] = 2
	fmt.Println(m2)

	//增、删 查 改
	//1 增 改 map添加元素在初始化时没定义的key即增加 当key存在时修改
	m1[3] = "history"
	m2["home"] = 3
	//改
	m1[1] = "music"
	fmt.Println(m1, m2)

	//查  遍历
	for k, v := range m1 {
		if v != "" {
			fmt.Println(k, v)
		}
	}
	fmt.Println("-------------------------")

	//删
	delete(m1, 1)
	delete(m2, "home")
	fmt.Println(m1, m2)
	fmt.Println("--------------------")

	//初始化key 为string value 为map的 map

	m3 := make(map[string]map[string]string)

	m3["bao"] = map[string]string{"add": "xm"}
	fmt.Println(m3)

	users := []string{"bao", "tae", "mei", "mei", "tae", "bao", "mei"}
	user_stat := make(map[string]int)

	for _, s1 := range users {
		if _, ok := user_stat[s1]; !ok {
			user_stat[s1] = 1
		} else {
			user_stat[s1]++
		}
	}
	fmt.Println(user_stat)

	m4 := make(map[int]string)
	m4[1] = "s"
	fmt.Println(m4)
	if _, k := m4[1]; !k {
		fmt.Println("错误")
	}
	fmt.Println(m4)

}
