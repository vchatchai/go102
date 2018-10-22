package sttime

import (
	"fmt"
	"time"
)

func Today() {
	today := time.Now()
	fmt.Println(today)
}
