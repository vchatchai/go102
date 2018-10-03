package stexec

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func WriteChileProcess() {
	cmd := []string{"sample"}
	proc := exec.Command(cmd[0])

	stdin, err := proc.StdinPipe()
	if err != nil {
		panic(err)
	}

	defer stdin.Close()

	stdout, err := proc.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()

	go func() {
		fmt.Println("output")
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:", s.Text())

		}
	}()

	stderr, err := proc.StderrPipe()
	if err != nil {
		panic(err)
	}
	stderr.Close()

	go func() {
		s := bufio.NewScanner(stderr)
		for s.Scan() {
			fmt.Println("Error", s.Text())
		}
	}()

	err = proc.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")
	time.Sleep(time.Second * 5)
	proc.Process.Kill()
}
