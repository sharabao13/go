// 字符串
package main

import (
	"fmt"
	"strings"
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	"unicode/utf8"
)

func main() {
	// `` 内的字符串无需转义
	path := `D:\tool\linux`
	// “” 内的字符串需要转义
	path1 := "\"D:\\tool\\linux\""
	name := "tae"
	email_Addr := "sharahong13@gmail.com"
	fmt.Println(path)
	fmt.Println(name)
	fmt.Println(path1)
	fmt.Println(email_Addr)
	// 字符串长度
	fmt.Println(len(email_Addr))
	// 字符串拼接
	s1 := "Hello"
	s2 := "World"
	s3 := s1 + s2
	// sprintf 不打印输出到终端，记录该值，要打印需复制变量
	s4 := fmt.Sprintf("%s%s\n", s1, s2)
	fmt.Printf("%s%s\n", s1, s2)
	fmt.Println(s1 + s2)
	fmt.Println(s3)
	fmt.Println(s4)
	//分隔 Split方法, strings.Split(变量，"分隔符")
	sp_1 := strings.Split(path, "\\")
	fmt.Println(sp_1)
	// 包含 Contains strings.Contains(变量，“值”),返回布尔值
	c_1 := strings.Contains(path, "tool")
	fmt.Println(c_1)
	// 前缀 后缀  HasPrefix HasSufix 方法 strings.HasPrefix(变量名,“值”)，返回布尔值
	s5 := "HelloWorld World"
	hP := strings.HasPrefix(s4, "Hello")
	hP1 := strings.HasPrefix(s4, "hello")
	hS := strings.HasSuffix(s5, "world")
	hS1 := strings.HasSuffix(s5, "World")
	fmt.Println(s4)
	fmt.Println(hP)
	fmt.Println(hP1)
	fmt.Println(hS)
	fmt.Println(hS1)
	// strings.Index 索引位置  strings.lastIndex 索引最后位置
	s6 := "aacdbbeadeffeaedbdeadeaddabeccdcadcaesssdeade"
	fmt.Println(strings.Index(s6, "c"))     //2
	fmt.Println(strings.LastIndex(s6, "c")) //34
	// 拼接 strings.Join 方法 strings.Join(变量,"符号")
	//s7 := "abcd"
	fmt.Println(strings.Join(sp_1, "+"))
	// rune ；类型
	ch_1 := "你好，中国"
	fmt.Println(len(ch_1))
	fmt.Printf("%T", ch_1)
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Printf("runeCountString:%T\n", utf8.RuneCountInString(ch_1))
	fmt.Println(utf8.RuneCountInString(ch_1))
	//通过rune类型处理unicode字符  []rune
	fmt.Printf("runeCount:%T\n", len([]rune(ch_1)))
	fmt.Println(len([]rune(ch_1)))
}
