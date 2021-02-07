package main

/*
	strings 包使用
	compare 比较字符串
	contains 包含字符串，包含一组字符串 containsany 包含任意一个
	counts 统计字符串出现次数
	EqualFold 不切分大小写比较
	Fields 空白字符切割 \t \n \r ..
	HasPrefix 判断是否某个字符串开头
	Index 判断字符串出现索引的位置 不存在返回-1
	LastIndex 最后一次出现的位置
	Split 按某种方式分割
	Join 按某种方式合并,[]string 切换
	Repeat  复制几份
	Replace 替换
	ToUpper 转换为大写
	ToLower 转换为小写


*/
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("cd", "ab"))
	fmt.Println(strings.Contains("abcdefg", "abcd"))
	fmt.Println(strings.ContainsAny("eladiefkds", "fdseadf"))
	fmt.Println(strings.Count("fdsafdsdfsaaaaaaafds", "a"))
	fmt.Println(strings.EqualFold("fdalfdsfsdf", "DFKEIALDFSA"))
	fmt.Println(strings.Fields("abc bb \ncc \tdd"))
	fmt.Println(strings.HasPrefix("abcdefg", "b"))
	fmt.Println(strings.HasSuffix("abcdefg", "d"))
	fmt.Println(strings.Index("aaabcdeeee", "bcd"))
	fmt.Println(strings.LastIndex("aaabbbcccdddaa", "aa"))
	fmt.Println(strings.Split("www.ibuy.ccb.com", "."))
	fmt.Println(strings.Join([]string{"www", "ibuy-demo", "ccb", "com"}, "."))
	fmt.Println(strings.Repeat("ab", 4))
	fmt.Println(strings.Replace("adcde", "ab", "12", 1))
	fmt.Println(strings.ToUpper("aaaaaaBBBB"))
	fmt.Println(strings.ToLower("BBBBBaaaaa"))
	fmt.Println(strings.Trim("xyzdkyzkzyz", "xz"))
}
