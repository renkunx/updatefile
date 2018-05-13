package main

import (
	"fmt"
)

const (
	USER     = "sybase"
	PASSWORD = "sybase123"
	HOST     = "192.168.216.129"
)

func main() {
	fmt.Println("Application  process ... ...")
	//建立 ssh 连接
	client := sshclientfunc(USER, PASSWORD, HOST)

	//使用 sftp 获取下载文件
	sftpGet("/home/sybase/bin", client)

	//读取配置，文件更新 mis.ini
	myConfig := new(Config)
	myConfig.InitConfig("mis.ini")
	fmt.Println(myConfig.Read("MISINFO", "COUNTER_NO"))
	fmt.Println("Application  end ... ...")
}
