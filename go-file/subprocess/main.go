/*
@Time : 2019-05-17 11:06 
@Author : xiaoyueya
@File : main
@Software: GoLand
*/
package main

import (
	"bufio"
"fmt"
"io"
"os"
"os/exec"
"time"

"syscall")

func main() {
	cmd1 := exec.Command("ping", "www.baidu.com")
	ppReader, err := cmd1.StdoutPipe()
	defer ppReader.Close()
	var bufReader = bufio.NewReader(ppReader)
	if err != nil {
		fmt.Printf("create cmd stdoutpipe failed,error:%s\n", err)
		os.Exit(1)
	}
	err = cmd1.Start()
	if err != nil {
		fmt.Printf("cannot start cmd1,error:%s\n", err)
		os.Exit(1)
	}
	go func() {
		var buffer []byte = make([]byte, 4096)
		for {
			n, err := bufReader.Read(buffer)
			if err != nil {
				if err == io.EOF {
				fmt.Printf("pipi has Closed\n")
				break
			} else {
				fmt.Println("Read content failed")
			}
			}
			fmt.Print(string(buffer[:n]))
		}
	}()
	time.Sleep(10 * time.Second)
	err = stopProcess(cmd1)
	if err != nil {
		fmt.Printf("stop child process failed,error:%s", err)
		os.Exit(1)
	}
	cmd1.Wait()
	time.Sleep(1 * time.Second)
}

func stopProcess(cmd *exec.Cmd) error {
	pro, err := os.FindProcess(cmd.Process.Pid)
	if err != nil {
		return err
	}
	err = pro.Signal(syscall.SIGINT)
	if err != nil {
		return err
	}
	fmt.Printf("结束子进程%s成功\n", cmd.Path)
	return nil
}
