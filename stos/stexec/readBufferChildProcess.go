package stexec

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ReadBufferChildProcess() {
	cmd := "ping"
	timeout := 2 * time.Second

	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "example.com")

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()
	proc.Start()

	s := bufio.NewScanner(stdout)

	for s.Scan() {
		fmt.Println(s.Text())
	}
}
