package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello updatefile project!")
	args := os.Args
	if len(args) < 2 {
		//不带参数，显示说明
		fmt.Println("这是一个依托 ftp 和 sftp 进行文件更新的程序")
		fmt.Println("                  Author By BluePrint")
		fmt.Println("Application  process ... ...")
	} else {
		//带参数
		fmt.Println(args)
	}

	// h := md5.New()
	// h.Write([]byte(args[1]))
	// fmt.Println(h.Sum(nil))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
