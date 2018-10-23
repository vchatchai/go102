package sttime

import (
	"fmt"
	"time"
)

func Unit() {
	t := time.Date(2018, 10, 22, 22, 30, 11, 0, time.Local)
	fmt.Printf("Extracting units from: %v\n", t)

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("The %dth day of %v is %v\n", dOfMonth, month, weekDay)
}
