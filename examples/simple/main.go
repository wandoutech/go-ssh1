package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ultram4rine/go-ssh1"
)

func main() {
	client, err := ssh1.Dial("localhost:2222", &ssh1.Config{
		CiphersOrder:    []int{ssh1.SSH_CIPHER_BLOWFISH, ssh1.SSH_CIPHER_3DES, ssh1.SSH_CIPHER_DES},
		User:            "root",
		AuthMethods:     []ssh1.AuthMethod{ssh1.Password("alpine")},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh1.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatal(err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	modes := ssh1.TerminalModes{
		ssh1.ECHO:          1,
		ssh1.TTY_OP_ISPEED: 14400,
		ssh1.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", 60, 80, modes); err != nil {
		log.Fatal(err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("in pipe", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal("out pipe", err)
	}

	session.Shell()

	var buf = make([]byte, 500)

	_, err = stdin.Write([]byte("mkdir test\n"))
	if err != nil {
		log.Fatal("can't write 1", err)
	}
	if _, err := stdout.Read(buf); err != nil {
		log.Println(err)
	}
	fmt.Print(string(buf))
	buf = nil

	_, err = stdin.Write([]byte("cd test\n"))
	if err != nil {
		log.Fatal("can't write 2", err)
	}
	if _, err := stdout.Read(buf); err != nil {
		log.Println(err)
	}
	fmt.Print(string(buf))
	buf = nil

	_, err = stdin.Write([]byte("touch file.txt\n"))
	if err != nil {
		log.Fatal("can't write 3", err)
	}
	if _, err := stdout.Read(buf); err != nil {
		log.Println(err)
	}
	fmt.Print(string(buf))
	buf = nil
}
