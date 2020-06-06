package main

import (
	"fmt"
)

func main() {

	/*
		map 跟 slice 结合使用
	*/
	map1 := make(map[string]string)
	map1["name"] = "li"
	map1["age"] = "20"
	map1["gender"] = "male"
	map1["addr"] = "quanzhou"
	fmt.Println(map1)

	map2 := map[string]string{"name": "Taylor", "age": "20", "gender": "fmale", "addr": "usa"}
	fmt.Println(map2)

	map3 := make(map[string]string)
	map3["name"] = "IU"
	map3["age"] = "18"
	map3["gender"] = "fmale"
	map3["addr"] = "korea"
	fmt.Println(map3)

	// 将map赋值给切片
	s1 := make([]map[string]string, 0, 3)
	//将map值添加进切片
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	fmt.Println(s1)

	for i, v := range s1 {
		//v map1 map2 map3
		fmt.Printf("第%d个的人的信息是: \n", i+1)
		fmt.Printf("\tname: %s\n", v["name"])
		fmt.Printf("\tage: %s\n", v["age"])
		fmt.Printf("\tgender: %s\n", v["gender"])
		fmt.Printf("\taddr: %s\n", v["addr"])
	}

}
