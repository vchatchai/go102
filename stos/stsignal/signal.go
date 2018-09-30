package stsignal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signal() {
	sChan := make(chan os.Signal, 1)

	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	exitChan := make(chan int)
	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CLRL+C")
			exitChan <- 1
		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1
		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()
	code := <-exitChan
	os.Exit(code)
}
