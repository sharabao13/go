package main

import "fmt"

// 创建一个人的结构体
type Human struct {
	name   string
	age    int
	weight int
}

//创建一个学生的字段
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	/*
		结构体匿名字段
		可以用字段来创建结构，这些字段只包含一个没有字段名的类型。这些字段被称为匿名字段。

		在类型中，使用不写字段名的方式，使用另一个类型
	*/

	// 初始化一个学生的信息
	mark := Student{Human{name: "a", age: 20, weight: 60}, "computer Science"}
	//访问相应的字段
	fmt.Println("his name is :", mark.name)
	fmt.Println("his age is :", mark.age)
	fmt.Println("his weight is :", mark.weight)
	fmt.Println("his is specilality is :", mark.speciality)

	//修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("mark change his speciality")
	fmt.Println("his is specilality is :", mark.speciality)

	/*
		可以使用"."的方式进行调用匿名字段中的属性值

		实际就是字段的继承

		其中可以将匿名字段理解为字段名和字段类型都是同一个

		基于上面的理解，所以可以mark.Human = Human{"Marcus", 55, 220} 和mark.Human.age -= 1

		若存在匿名字段中的字段与非匿名字段名字相同，则最外层的优先访问，就近原则

		通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。
	*/

}
