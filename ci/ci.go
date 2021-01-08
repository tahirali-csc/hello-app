package main

import (
	"bufio"
	"log"
	"os/exec"

	ci "github.com/tahirali-csc/task-executor-ci"
)

func runCommand(cmd string) (string, error) {

	cmdObj := exec.Command("/bin/sh", "-c", cmd)
	reader, err := cmdObj.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmdObj.Start()
	scanner := bufio.NewScanner(reader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		done <- true
	}()

	<-done
	err = cmdObj.Wait()
	return "", err
	// return string(out), err
}

func main() {
	log.Println("Starting Build...")

	//For testing
	//os.Setenv("TE_HOST_URL", "http://localhost:8080")
	//os.Setenv("TE_BUILD_ID", "17")

	//log.Println("Running Pipeline")
	//output, err := runCommand("cd ../ && go test -v -run TestRun1 hello-app/pkg/module1")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(output)
	//
	//output, err = runCommand("cd ../ && go test -v -run TestRun2 hello-app/pkg/module1")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(output)

	build := ci.NewBuild()

	err := build.Exec(&ci.Step{
		Name:  "unit-test",
		Image: "alpine:latest",
		Cmd:   []string{"/bin/sh", "-c", "date"},
	})
	if err != nil {
		log.Println(err)
	}

	err = build.Exec(&ci.Step{
		Name:  "int-test",
		Image: "alpine:latest",
		Cmd:   []string{"/bin/sh", "-c", "ls -al / && echo 'Sleeping' && sleep 10s && pwd"},
	})
	if err != nil {
		log.Println(err)
	}

	build.Done()
}
