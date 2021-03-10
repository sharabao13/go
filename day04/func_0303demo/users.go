package main

import (
	"fmt"
	"strings"

	//"go/ast"
	"strconv"
)

const (
	passWord = "123abc"
	maxAuth  = 3
)

var passWordInfo string

//用户认证
func auth() bool {
	for i := 0; i < maxAuth; i++ {
		fmt.Println("请输入密码: ")
		fmt.Scan(&passWordInfo)
		if passWordInfo == passWord {
			fmt.Println("登录成功")
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

//添加用户函数
func addUser(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Println("请输入姓名: ")
	fmt.Scan(&name)
	fmt.Println("请输入年龄: ")
	fmt.Scan(&age)
	fmt.Println("请输入联系方式: ")
	fmt.Scan(&tel)
	fmt.Println("请输入住址: ")
	fmt.Scan(&addr)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}

}

//用户查询
func query(users map[string]map[string]string) {
	var inputInfo string
	fmt.Println("请输入查询信息")
	fmt.Scan(&inputInfo)
	title := fmt.Sprintf("%5s|%25s|%5s|%25s|%30s", "Id", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	for _, user := range users {
		if strings.Contains(user["name"], inputInfo) || strings.Contains(user["age"], inputInfo) || strings.Contains(user["tel"], inputInfo) || strings.Contains(user["addr"], inputInfo) {
			fmt.Printf("%5s|%25s|%5s|%25s|%30s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}

//用户修改
func modify(users map[string]map[string]string) {
	var inputId string
	fmt.Println("请输入要修改的用户ID: ")
	fmt.Scan(&inputId)
	for _, user := range users {
		if user["id"] != inputId {
			fmt.Println("用户不存在或输入错误")
		} else {
			fmt.Printf("%5s|%25s|%5s|%25s|%30s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}
func main() {
	//命令行管理系统
	/*
		    用户登录
				在程序中定义password = “123abc”
				提示输入密码，如果密码输入3次错误，提示并退出
				3次内密码验证成功继续操作
			用户信息存储
				--> 内存  数据存放
				--> 结构  数据类型map
				--> 用户	  id name age tel addr
		                  map 数据结构
		                  使用的数据类型暂时使用string
		    用户添加
			用户删除
					输入要修改的用户ID:
						id 在在或不存在输入错误
						打印用户信息，提示用户是否确认删除
						删除或退出
			用户查询
			用户修改
				输入要修改的用户ID:
					id 在在或不存在输入错误
					打印用户信息，提示用户是否确认修改
					修改 name age tel addr
	*/
	//定义一个用户

	users := make(map[string]map[string]string)
	id := 0
	if !auth() {
		fmt.Printf("密码输入%d次错误,程序退出\n", maxAuth)
		return
	}
	fmt.Println("欢迎使用xx用户管理系统")
	//fmt.Println("请输入密码: ")
	//fmt.Scan(&passWordInfo)
	// 不是函数的认证方法	//for {
	//	//	var count int
	//	//	inputPass:
	//	//		fmt.Println("请输入密码: ")
	//	//		fmt.Scanln(&passWordInfo)
	//	//		if passWordInfo == passWord {
	//	//			fmt.Println("登录成功")
	//	//			break
	//	//		}
	//	//		fmt.Println("密码输入错误")
	//	//		count ++
	//	//	    if count > 3 {
	//	//			panic("密码错误次数够多,,main...over..")
	//	//		}
	//	//		goto inputPass
	//	//}

	for {
		var op string
		fmt.Print(`
		1.新增用户
		2.修改用户
		3.查询用户
		4.删除用户
		5.退出
		请输入操作指令: 
	`)
		fmt.Scan(&op)

		fmt.Println(op)

		switch op {
		case "1":
			id++
			addUser(id, users)
		case "2":
			modify(users)
		case "3":
			query(users)
		case "4":
			query(users)
		case "5":
			goto QUIT
		default:
			fmt.Println("输入有误，请重新输入")
		}

		//fmt.Println("你输入的指令为: ",op)
	}
QUIT:
	fmt.Println("退出系统")
}
