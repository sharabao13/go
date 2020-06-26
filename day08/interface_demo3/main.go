package main

import "fmt"

func main() {
	/*
		空接口interface()
			不包含任何的方法，，正因为如此，所有类型都实现了空接口，因此空接口可以实现存储任意类型的数值
	*/

	var a1 A = Cat{collor: "白色"}
	var a2 A = Person{name: "a", age: 20}
	var a3 A = "b"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	//调用空接口函数
	test(a1)
	test(a2)
	test(3.13)
	test("cc")

	test1(a3)
	test1(a4)
	test1(5.56)
	test1("ruby")

	//map key字符串 value任意数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "tae"
	map1["age"] = 30
	map1["friend"] = Person{name: "jerry", age: 20}
	fmt.Println(map1)

	//切片 存储任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, "mei", 100)
	fmt.Println(slice1)

	test3(slice1)
}

//变量切片函数
func test3(slice []interface{}) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("第%d个数据为: %v\n", i+1, slice[i])
	}

}

//定义一个函数，参数为空接口类型，接收任意数据类型的参数
func test(a A) {
	fmt.Println(a)
}

//定义一个函数，参数为匿名空接口类型，接收任意数据类型的参数
func test1(a interface{}) {
	fmt.Println(a)
}

//定义一个空接口
type A interface {
}

type Cat struct {
	collor string
}

type Person struct {
	name string
	age  int
}
