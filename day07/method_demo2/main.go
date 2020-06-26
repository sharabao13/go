package main

import (
	"fmt"
)

func main() {
	/*
		oop 继承性
			如果两个类存在继承性，其中一个是子类，另一个作为父类，那么，
				1.子类可以访问父类的方法
				2.子类可以新增自己的熟悉和方法
				3.子类可以重写父类的方法（override 就是将父类已有的方法，重新实现）

		go 语言的结构嵌套
			1.模拟继承性 is -a
				type a struct{
					field
			}
				type b struct{
					a //匿名字段
			}

			2.模拟聚合关系 has -a

				type c struct{
					field
			}

				type d struct{
					c c //聚合关系
			}
	*/

	//1.创建persion对象
	p1 := Person{name: "a", age: 20}
	fmt.Println(p1.name, p1.age) //访问父类的字段属性
	p1.eat()                     //父类对象 调用父类的方法

	//2. 创建student对象
	p2 := Student{Person{name: "b", age: 20}, "qinghua"}
	fmt.Println(p2.name)
	fmt.Println(p2.age)    //子类对象，可以访问父类的字段，提升字段
	fmt.Println(p2.School) //子类对象，可以访问自己新增的字段属性
	p2.eat()               //子类对象，可以访问父类的方法

	p2.study() //子类对象方法自己新增的方法
	p2.eat()   //子类对象访问重写的父类方法
}

//定义一个"父类"

type Person struct {
	name string
	age  int
}

//定义一个"子类"
type Student struct {
	Person //结构体嵌套，模拟继承性
	School string
}

// 定义一个父类的方法
func (p Person) eat() {
	fmt.Println("父类的方法，吃饭")
}

//子类新增自己的方法

func (s Student) study() {
	fmt.Println("子类新增方法，学习")
}

//子类重写父类的方法
func (s Student) eat() {
	fmt.Println("子类重写方法，吃炸鸡")
}
