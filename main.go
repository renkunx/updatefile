package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("hello updatefile project!")
	// args := os.Args
	// if len(args) < 2 {
	// 	//不带参数，显示说明
	// 	usage()
	// } else {
	// 	//带参数
	// 	usage()
	// 	fmt.Println(args)
	// }

	// h := md5.New()
	// h.Write([]byte(args[1]))
	// fmt.Println(h.Sum(nil))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))

	//测试二维码生成
	genqrcodepng("qrcode.png", "showsomethig", qrcode.Medium)

	//测 ftp
	// SSH("sybase", "sybase123", "192.168.216.129:22")
	// ftpdemo()
	sshclientfunc()
}

func usage() {
	fmt.Println("这是一个依托 ftp 和 sftp 进行文件更新的程序")
	fmt.Println("                  Author By BluePrint")
	fmt.Println("Application  process ... ...")
}
