// 调用shell命令在当前路径下创建目录
package main

// 调用shell 命令 "os/exec" 模块
import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("touch", "test_file")

	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command Failed:" + err.Error())
		return
	}
	fmt.Println("Execute Command Success.")
}
