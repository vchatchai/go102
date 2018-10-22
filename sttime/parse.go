package sttime

import (
	"fmt"
	"time"
)

func Parse() {
	t, err := time.Parse("2/1/2006", "31/7/2018")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	t, err = time.ParseInLocation("2/1/2006 3:04 PM", "31/7/2018 11:48 AM", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
