package stexec

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func ReadStdOut() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)

	buf := bytes.NewBuffer([]byte{})

	proc.Stdout = buf
	proc.Run()
	fmt.Println(string(buf.Bytes()))
}
