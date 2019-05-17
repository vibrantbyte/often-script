/*
@Time : 2019-05-17 11:39 
@Author : xiaoyueya
@File : mutilprocess
@Software: GoLand
*/
package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

func main() {

	cmd := exec.Command("./mutilprocess")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("nobody")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		uid, _ := strconv.Atoi(user.Uid)
		gid, _ := strconv.Atoi(user.Gid)

		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
