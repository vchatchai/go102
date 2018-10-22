package sttime

import (
	"fmt"
	"time"
)

func Epoch() {
	t := time.Unix(0, 0)
	fmt.Println(t)

	epoch := t.Unix()
	fmt.Println(epoch)

	apochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seond: %d\n", apochNow)

	apochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-second: %d\n", apochNano)
}
