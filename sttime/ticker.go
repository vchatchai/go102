package sttime

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Ticker() {
	c := make(chan os.Signal, 1)

	signal.Notify(c)

	ticker := time.NewTicker(time.Second)

	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick", t)
			case <-stop:
				fmt.Println("GOrountine closing")
				return
			}
		}
	}()

	<-c
	ticker.Stop()

	stop <- true

	fmt.Println("Application stopped")
}
