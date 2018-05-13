package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"

	"fmt"
	"github.com/pkg/sftp"
	"io"
	"time"
)

func sftpGet(rometdir string, client *ssh.Client) {
	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()
	// walk a directory
	w := sftp.Walk(rometdir)

	// 追加更新日志
	log.Printf(rometdir + "/updatefile.log")
	logfile, err := sftp.OpenFile(rometdir+"/updatefile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY)
	defer logfile.Close()
	if _, err := logfile.Write([]byte(fmt.Sprintf("%s updating... \n", gethostip()))); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		fi, err := sftp.Lstat(w.Path())
		if err != nil {
			log.Fatal(err)
		}
		//过滤文件夹和更新日志文件
		if !fi.IsDir() && fi.Name() != "updatefile.log" {
			remoteFile, err := sftp.Open(w.Path())
			if err != nil {
				log.Fatal(err)
			}
			defer remoteFile.Close()

			localFile, err := os.Create(getCurrentPath() + fi.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer localFile.Close()
			size := fi.Size()
			start := time.Now()
			n, err := io.Copy(localFile, io.LimitReader(remoteFile, size))
			if err != nil {
				log.Fatal(err)
			}
			if n != size {
				log.Fatalf("copy: expected %v bytes, got %d", size, n)
			}
			log.Printf("'%s' read %v bytes in %s", fi.Name(), size, time.Since(start))

			if _, err := logfile.Write([]byte(fmt.Sprintf("'%s' read %v bytes in %s \n", fi.Name(), size, time.Since(start)))); err != nil {
				log.Fatal(err)
			}
		}
	}

}
