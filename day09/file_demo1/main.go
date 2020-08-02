package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	/*
		文件操作
			1.相对路径 relative
				ab.txt

			2.绝对路径 absolute
				E:\\untitled\\go_ducment\a.txt
			 当前目录 .
			 上一层目录 ..
			filepath
	*/
	fileName1 := "ab.txt"
	fileName2 := `E:\\untitled\\go_ducment\a.txt`
	// 判断是否绝对路径路径 filepath.Isabs
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))

	//打印文件绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	//获取父目录 文件只能用绝对路径
	fmt.Println(filepath.Join(fileName2, ".."))

	////创建目录 os.mkdir //创建单个目录
	//err := os.Mkdir(`E:\\untitled\\go_ducment\b`,os.ModePerm)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("目录创建成功")
	////os.mkdirall 创建多层目录包含子目录
	//err1 := os.MkdirAll(`E:\\untitled\\go_ducment\c\a\b`,os.ModePerm)
	//if err != nil{
	//	fmt.Println(err1)
	//}
	//fmt.Println("多层目录创建成功")

	//创建文件 os.create 默认模式0666 任何人可读可写，如果文件存在会截断他（空文件
	//file1,err2 := os.Create(`E:\\untitled\\go_ducment\c.txt`)
	//if err2 != nil{
	//	fmt.Println("err",err2)
	//	return
	//}
	//fmt.Println(file1)

	// 打开文件 os.open 默认只读模式
	//fileName3 := `E:\\untitled\\go_ducment\a.txt`
	//file2,err3 := os.Open(fileName3)
	//if err3 != nil{
	//	fmt.Println("err",err3)
	//	return
	//}
	//fmt.Println(file2)

	//打开并操作文件os.openfile
	/*
		第一个参数 文件名
		第二个参数 文件打开方式
			// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
		第三个 文件的权限 如果文件不存在 创建文件
	*/
	//file3,err4 := os.OpenFile(fileName3,os.O_RDWR,os.ModePerm)
	//if err4 != nil{
	//	fmt.Println("err",err4)
	//	return
	//}
	//fmt.Println(file3)
	//关闭程序和文件链接 close
	//file3.Close()

	//删除目录或文件os.remove  删除包含子目录os.removeall
	//fileName3 := `E:\\untitled\\go_ducment\b`
	//err := os.Remove(fileName3)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("删除文件成功")

	fileName3 := `E:\\untitled\\go_ducment\c`
	err := os.RemoveAll(fileName3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除文件成功")

}
