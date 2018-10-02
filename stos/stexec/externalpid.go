package stexec

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func ExternalPid() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "1")

	proc.Start()

	proc.Wait()
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited sucessfuly: %t\n", proc.ProcessState.Success())
}
