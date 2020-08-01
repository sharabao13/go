package main

import "fmt"

func main() {
	lenth, width := -5.67, -8.72
	area, err := rectArea(lenth, width)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*rectError); ok {
			if err.widthNative() {
				fmt.Println(err.width)
			}
			if err.lenthNative() {
				fmt.Println(err.lenth)
			}
		}
		return
	}
	fmt.Printf("长方形的面积是: %.2f", area)
}

// 定义一个结构体表示错误信息
type rectError struct {
	msg          string  //错误信息描述
	lenth, width float64 //错误的长度，宽度
}

// 定义结构体msg的方法
func (e *rectError) Error() string {
	return e.msg
}

//错误长度的方法
func (e *rectError) lenthNative() bool {
	return e.lenth < 0
}

//错误宽度的方法
func (e *rectError) widthNative() bool {
	return e.width < 0
}

func rectArea(lenth, width float64) (float64, error) {
	msg := ""
	if lenth < 0 {
		msg = "长度小于0"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度小于0"
		} else {
			msg += ",宽度也小于0"
		}
	}
	if msg != "" {
		return 0, &rectError{msg, lenth, width}
	}
	return lenth * width, nil
}
