package sttime

import (
	"fmt"
	"time"
)

func Format() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)
	fmt.Println("tTime:", tTime)
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	fmt.Printf("the time is: %s\n", tTime.Format("15:04"))
	fmt.Printf("the time is: %s\n", tTime.Format(time.RFC1123))
	fmt.Printf("the time is: %s\n", tTime.Format("2006/01/_2"))
	fmt.Printf("the time is: %s\n", tTime.Format("2006/01/02"))
	fmt.Printf("the time is: %s\n", tTime.Format("15:04:05.00"))
	fmt.Printf("the time is: %s\n", tTime.Format("15:04:05.999"))
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up:"), "03:04PM")))
}
