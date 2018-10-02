package stexec

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OutputChildProcess() {
	var cmd string
	if runtime.GOOS == "window" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)
	buff, err := proc.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buff))
}
