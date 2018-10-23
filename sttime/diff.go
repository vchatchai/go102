package sttime

import (
	"fmt"
	"time"
)

func Diff() {
	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2018, 10, 22, 22, 54, 1, 1, l)
	t2 := time.Date(2017, 9, 21, 22, 54, 1, 1, l)

	fmt.Printf("First Default date  is %v\n", t)
	fmt.Printf("Second Default date is %v\n", t2)

	dur := t.Sub(t2)
	fmt.Printf("The duration between t and t2 is %v\n", dur)

	dur = time.Since(t)
	fmt.Printf("The duration between now and t is %v\n", dur)

	dur = time.Until(t)
	fmt.Printf("The duration between t and now is %v\n", dur)
}
