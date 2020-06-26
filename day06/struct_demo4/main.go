package main

import "fmt"

func main() {
	/*
		结构体嵌套
			一个结构体可能包含一个字段，而这个字段反过来就是一个结构体，这些结构体被称为嵌套结构
	*/

	var p Person
	p.name = "Navven"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illions",
	}
	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
	fmt.Println("city", p.address.city)
	fmt.Println("state", p.address.state)

	b1 := Book{
		"xiyouji",
		26.5,
	}
	s1 := Students{
		"a",
		20,
		b1,
	}
	fmt.Println(b1)
	fmt.Println(s1)
	fmt.Printf("姓名: %s,年龄: %d,书名: %s,价格:%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)

	s1.book.bookName = "xxxxxxxx"
	fmt.Println(s1)
	fmt.Println(b1)

	s2 := Students{"b", 18, Book{"hongloumeng", 78.89}}
	fmt.Println(s2)

	s3 := Students{
		name: "c",
		age:  20,
		book: Book{
			bookName: "shuihuzhuang",
			price:    78.69,
		},
	}
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookName, s3.book.price)

	s4 := Students2{"d", 18, &Book{"AA", 88.88}}
	fmt.Println(s4)
	fmt.Println(*s4.book)

	s4.book.bookName = "nnnnnn"
	fmt.Println(s4)
	fmt.Println(*s4.book)
}

type Address struct {
	city, state string
}

type Person struct {
	name    string
	age     int
	address Address
}

//定义一个书的结构体
type Book struct {
	bookName string
	price    float64
}

//定义一个学生的结构体
type Students struct {
	name string
	age  int
	book Book
}

type Students2 struct {
	name string
	age  int
	book *Book
}
