package stsignal

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func Shutdown() {
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log", time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	closeChan := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		select {
		case <-closeChan:
			log.Println("Goroutine closing")
			return
		default:
			log.Println("Writing to log")
			io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
	)
	<-sigChan

	close(closeChan)

	releaseAllResource()

	fmt.Println("The application shut down gracefully")
}

func releaseAllResource() {
	io.WriteString(writer, "Application releasing all resource")
	writer.Close()
}
