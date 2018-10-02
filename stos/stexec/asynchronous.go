package stexec

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Asynchronous() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(prc.Stdout)
	}
}
