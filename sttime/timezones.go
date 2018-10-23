package sttime

import (
	"fmt"
	"time"
)

func TimeZones() {
	asia, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	t := time.Date(2018, 10, 22, 23, 36, 0, 0, asia)

	fmt.Printf("Orginal Time: %v\n", t)

	phx, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		panic(err)
	}

	t2 := t.In(phx)
	fmt.Printf("Converted Time: %v\n", t2)
}
