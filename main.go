package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("hello updatefile project!")
	args := os.Args
	if len(args) < 2 {
		//不带参数，显示说明
		usage()
	} else {
		//带参数
		usage()
		fmt.Println(args)
	}
	// SSH("sybase", "sybase123", "192.168.216.129:22")
	ftpdemo()
	// h := md5.New()
	// h.Write([]byte(args[1]))
	// fmt.Println(h.Sum(nil))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

func usage() {
	fmt.Println("这是一个依托 ftp 和 sftp 进行文件更新的程序")
	fmt.Println("                  Author By BluePrint")
	fmt.Println("Application  process ... ...")
}

func ftpdemo() {
	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: "mon",
		Auth: []ssh.AuthMethod{
			ssh.Password("mon123"),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "192.168.216.129:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sftp.Getwd())

	defer sftp.Close()

	// walk a directory
	w := sftp.Walk("/home/sybase")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	// leave your mark
	f, err := sftp.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}

	// check it's there
	fi, err := sftp.Lstat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)
}

// func SSH(user, password, ip_port string) {
// 	PassWd := []ssh.AuthMethod{ssh.Password(password)}
// 	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
// 	Client, _ := ssh.Dial("tcp", ip_port, &Conf)
// 	defer Client.Close()
// 	a := bufio.NewReader(os.Stdin)
// 	for {
// 		b, _, z := a.ReadLine()
// 		if z != nil {
// 			return
// 		}
// 		command := string(b)
// 		if session, err := Client.NewSession(); err == nil {
// 			defer session.Close()
// 			session.Stdout = os.Stdout
// 			session.Stderr = os.Stderr
// 			session.Run(command)
// 		}
// 	}
// }
