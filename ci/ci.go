package main

import (
	"log"
	"os/exec"
)

func runCommand(cmd string) (string, error) {
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	return string(out), err
}

func main() {
	log.Println("Running Pipeline")
	output, err := runCommand("cd ../ && go test -v -run TestRun1 hello-app/pkg/module1")
	if err != nil {
		panic(err)
	}
	log.Println(output)

	output, err = runCommand("cd ../ && go test -v -run TestRun2 hello-app/pkg/module1")
	if err != nil {
		panic(err)
	}
	log.Println(output)
}
