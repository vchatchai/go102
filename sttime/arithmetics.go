package sttime

import (
	"fmt"
	"time"
)

func Arithmetics() {
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	t := time.Date(2018, 10, 22, 22, 39, 01, 10, l)
	fmt.Printf("Default date is: %v\n", t)

	//add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date +3HRS is: %v\n", r1)

	//substract 3 days
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3HRS is :%v\n", r1)

	//add day/month/year
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1YR +3MTH +2D is: %v\n", r1)
}
