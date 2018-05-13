package main

import (
	"os/exec"
	"os"
	"path/filepath"
)

func getCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	currentfile, _ := filepath.Abs(file)
	dir, file := filepath.Split(currentfile)
	return dir
}

// args := os.Args
// if len(args) < 2 {
// 	//不带参数，显示说明
// 	usage()
// } else {
// 	//带参数
// 	usage()
// 	fmt.Println(args)
// }
//测试 md5
// h := md5.New()
// h.Write([]byte(args[1]))
// fmt.Println(h.Sum(nil))
// fmt.Println(hex.EncodeToString(h.Sum(nil)))

//测试二维码生成
//genqrcodepng("qrcode.png", "showsomethig", qrcode.Medium)