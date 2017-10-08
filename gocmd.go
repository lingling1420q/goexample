package main

import (
	"bytes"
	//"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func Cmd(cmds string) (result string) {
	log.Println("run cmd:", cmds)
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash", "-c", cmds)
	var domifstat bytes.Buffer
	cmd.Stdout = &domifstat
	err := cmd.Run()
	if err != nil {
		log.Printf("Error while exec cmd %a", err)
		return ""
	}
	result = domifstat.String()
	return
}

func CmdAsync(cmds string) (result string) {
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash", "-c", cmds)
	stdout, err := cmd.StderrPipe()
	if err != nil {
		log.Println("error:", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// b, _ := ioutil.ReadAll(stdout)
	// log.Println(string(b))
	buffer := make([]byte, 1024)
	for {
		n, err := stdout.Read(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("n:", n, string(buffer))
		if n == 0 {
			time.Sleep(1 * time.Second)
		}
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	// result := CmdAsync("ffmpeg -i  /root/right.mp4 -b:a 500k -y /root/right_out.mp4")
	// log.Println(result)
	err := exec.Command("/bin/sh", "-c", "title", "yangaowei").Start()
	log.Println(err)
}
