package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		前面说过，因为空接口 interface{}没有定义任何函数，因此 Go 中所有类型都实现了空接口。当一个函数的形参是interface{}，
		那么在函数中，需要对形参进行断言，从而得到它的真实类型。

		// 安全类型断言

		<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 )

		//非安全类型断言

		<目标类型的值> := <表达式>.( 目标类型 )

		接口断言：
			方式一 ：
				1.interface := 接口对象.(实际类型)不安全 会panic（）

				2.interface,ok := 接口对象.(实际类型) 安全

			方式二：
				switch
				switch interface := 接口对象.(type){
					case实际对象1:
						...
					case实际对象2：
						...
					...
				}
	*/

	var t1 Triangle = Triangle{a: 3, b: 4, c: 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)
	fmt.Println("------------------")

	var c1 Circle = Circle{radius: 4}
	fmt.Println(c1.area())
	fmt.Println(c1.peri())
	fmt.Println(c1.radius)
	fmt.Println("------------------")

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())
	fmt.Println("------------------")

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())
	fmt.Println("------------------")

	testShape(t1)
	testShape(c1)
	testShape(s1)
	testShape(s2)
	fmt.Println("------------------")

	getType(t1)
	getType(c1)
	getType(s1)
	getType(s2)
	fmt.Println("--------------------")

	var t2 *Triangle = &Triangle{a: 2, b: 4, c: 2}
	fmt.Printf("%T,%p\n", t2, &t2)
	getType(t2)
	fmt.Println("-----------------")
	getType2(t2)
	getType2(c1)
}
func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针", ins.a, ins.b, ins.c)
	}
}

func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是:", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是: ", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("%T,%p\n", ins, &ins)
		fmt.Printf("%T,%p\n", s, &s)
	} else {
		fmt.Println("nil")
	}
}

// 定义一个接口函数
func testShape(s Shape) {
	fmt.Printf("周长: %.2f,面积: %.2f\n", s.peri(), s.area())
}

//定义一个 形状接口
type Shape interface {
	peri() float64 //形状的周长
	area() float64 //形状的面积
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
