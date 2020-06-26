package main

import "fmt"

//创建一个书的结构体
type BookS struct {
	title   string
	author  string
	subject string
	book_id int
}

//创建一个打印数函数
func printBook(book BookS) {
	fmt.Printf("book titile: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
	fmt.Printf("book subject: %s\n", book.subject)
	fmt.Printf("book book_id: %d\n", book.book_id)
}

func main() {
	/*
		结构体作为函数的参数
	*/
	var Book1 BookS //声明book1 为BookS 类型
	var Book2 BookS //声明book1 为BookS 类型

	/* book1 的描述 */
	Book1.title = "GO语言"
	Book1.author = "Bao"
	Book1.subject = "GO"
	Book1.book_id = 1

	/* book2 的描述 */
	Book2.title = "Pythoin语言"
	Book2.author = "Bao"
	Book2.subject = "Python"
	Book2.book_id = 2

	/* 打印book1的信息 */
	printBook(Book1)

	/* 打印book2的信息 */
	printBook(Book2)
}
