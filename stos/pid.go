package stos

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func PID() {
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)
	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}
