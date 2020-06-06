package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		strings包下关于字符串函数的使用
		常用函数使用

	*/

	s1 := "Hello World"
	// func Contains(s, substr string) bool
	// strings.Contains 方法匹配 substr 是否包含在s 返回bool 区分大小写
	fmt.Println(strings.Contains(s1, "hell"))

	//func ContainsAny(s, chars string) bool
	// strings.ContainsAny 方法 匹配substr内的字符任意一个字符匹配返回true， 不区分字符大小写
	fmt.Println(strings.ContainsAny(s1, "qqqH"))

	//func Count(s, substr string) int
	// strings.Count 方法 统计substr 在s内出现的次数 返回int substr 是一起匹配的
	fmt.Println(strings.Count(s1, "l"))

	// func HasPrefix(s, prefix string) bool
	// strings.HashPrefix 左前缀开始匹配 返回bool
	s2 := "202006strings.txt"
	if strings.HasPrefix(s2, "202005") {
		fmt.Println("文档是202005月份的")
	}
	//func HasSuffix(s, suffix string) bool
	//strings.HasSuffix 有前缀匹配 返回bool
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("txt文档格式")
	}

	//func Index(s, substr string) int
	//strings.Index 查询substr的索引下标,如不存在返回-1,返回第一次出现的下标
	fmt.Println(strings.Index(s1, "lo"))

	//func IndexAny(s, chars string) int
	//strings.IndexAny 查询substr的索引下标，返回从左开始找到下标，返回一次，不存在返回-1
	fmt.Println(strings.IndexAny(s1, "ld"))

	//func LastIndex(s, substr string) int
	//strings.LastIndex 从最后开始查询匹配substr的下标，返回1次，不存在返回-1
	fmt.Println(strings.LastIndex(s1, "ld"))

	//func Join(elems []string, sep string) string
	//strings.Join字符串拼接 对一组切片的元素进行拼接，sep：按什么符号拼接 返回一组新增的字符串需重新赋值
	Slice1 := []string{"我", "string", "89", "fmale"}
	s3 := strings.Join(Slice1, "/")
	fmt.Println(s3)
	// 将www youtube com 按.进行拼接
	netWorkAddr := []string{"www", "youtube", "com"}
	newNetWorkAddr := strings.Join(netWorkAddr, ".")
	fmt.Println(newNetWorkAddr)

	//func Split(s, sep string) []string
	//strings.Split 分隔 对字符串进行分隔 sep：按什么符号进行分隔 返回新的字符需重新赋值

	emailAddr := "sharehong13@gmail.com"
	sEmailAddr := strings.Split(emailAddr, "@")
	fmt.Println(sEmailAddr)

	//func SplitAfterN(s, sep string, n int) []string
	//strings.SplitAfterN 方法 从后n之前的的位置开始切分 n后面的保留一组字符串
	//s4 := "a,b,c,hello,world,ssj"
	//ns4 := strings.SplitAfterN(s4, ",", 1)
	//fmt.Println(ns4)

	//fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))

	//func Repeat(s string, count int) string
	//strings.Repeat 按字符串重复拼接n次

	s5 := strings.Repeat("abc", 3)
	fmt.Println(s5)

	//func Replace(s, old, new string, n int) string 替换
	//strings.Replace 将字符串内的字符进行替换，n代表替换的次数 n如果是-1 代表全部替换

	s6 := strings.Replace(s1, "l", "o", 4)
	fmt.Println(s6)

	//func ToLower(s string) string
	//strings.ToLower 将大写字符全部转为小写 特殊符号保持变
	s7 := "Hello Wrold,sse,333.***"
	fmt.Println(strings.ToLower(s7))

	//func ToUpper(s string) string
	//strings.ToUpper 将小写字符全部转为小写 特殊符号保持变
	fmt.Println(strings.ToUpper(s7))

	// 字符串截取子串
	//str[start:end]
	fmt.Println(s1[:5])

}
