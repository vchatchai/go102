package sttime

import (
	"fmt"
	"time"
)

func TimeOut() {
	to := time.After(3 * time.Second)

	list := make([]string, 0)
	done := make(chan bool, 1)

	fmt.Printf("Starting to insert items")
	go func() {
		defer fmt.Println("Exiting goroutine")
		for {
			select {
			case <-to:
				fmt.Println("The time is up")
				done <- true
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	fmt.Printf("Managed to insert %d items\n", len(list))
	fmt.Println(list[0:3])
}
