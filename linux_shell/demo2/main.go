package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	//cmd := exec.Command("touch", "test_file")
	//cmd1 := exec.Command()
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Execute Command failed:" + err.Error())
	//	return
	//}
	//
	//fmt.Println("Execute Command finished.")

	//mkdirFiles := exec.Command("mkdir","otter")
	//err := mkdirFiles.Run()
	//if err != nil {
	//	fmt.Println("命令执行失败：" + err.Error())
	//	return
	//}
	//fmt.Println("命令执行完成")

	//command := `./dir_size.sh .`
	//cmd := exec.Command("/bin/bash", "-c", command)
	//
	//output, err := cmd.Output()
	//if err != nil {
	//	fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
	//	return
	//}
	//fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

	cmd := exec.Command("/bin/bash", "-c", "grep test")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	stdin.Write([]byte("go text for grep\n"))
	stdin.Write([]byte("go test text for grep\n"))
	stdin.Close()

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:" + string(out_bytes))

}
